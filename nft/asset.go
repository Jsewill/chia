package rpc

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Asset represents an NFT asset
type Asset struct {
	Uris []string
	hash string
}

// Hash retrieves, or computes and compares, the hash for this Asset from its URI(s), and return the hash if they agree, otherwise return an error.
func (a *Asset) Hash() (string, error) {
	// If hash is set, don't compute.
	if a.hash != "" {
		return a.hash, nil
	}
	if len(a.Uris) == 0 {
		err := fmt.Errorf("There are no URLs to hash on this Asset.")
		logErr.Println(err)
		return "", err
	}
	hashMap := make(map[string]string)
	var prevHash string
	// Get hashes for all URLs and compare.
	for i, u := range a.Uris {
		s := sha256.New()
		// Do we have a legal URL, or is it possibly a file path?
		url, err := url.Parse(u)
		if err == nil && u.Hostname() != "" {
			// Retrieve asset from URL.
			r, err := http.Get(url.String())
			if err != nil {
				// Couldn't get asset from URL
				err = fmt.Errorf("Unable to get asset at %s for hashing.", u)
				logErr.Println(err)
				return "", err
			}
			defer r.Body.Close()
			// Attempt to get the hash of the asset.
			_, err := io.Copy(s, r.Body)
		} else {
			// We may have a file path. Error handling comes later.
			afile, err := os.Open(u)
			// Attempt to get the hash of the asset.
			_, err = io.Copy(s, afile)
		}
		// Handle hashing error.
		if err != nil {
			err = fmt.Errorf("Couldn't read from asset at %s, to get its hash.", u)
			logErr.Println(err)
			return "", err
		}
		// Hash and check against the previous hash.
		hashMap[u] = hex.EncodeToString(s.Sum(nil))
		if i > 0 && hashMap[u] != prevHash {
			err = fmt.Errorf("Hash of asset at %s, is not identical to one of the others: %s.", u, prevHash)
			logErr.Println(err)
			return "", err
		}
		prevHash = hashMap[u]
	}
	// Set hash, having successfully hashing each URL, and checking for duplicates.
	a.hash = hashMap[a.Uris[0]]

	return a.hash, nil
}

// SetHash sets the asset hash to the supplied string.
func (a *Asset) SetHash(h string) {
	a.hash = h
}
