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
	Hash string
}

// Retrieve and compare hashes for this Asset from its URI(s)
func (a *Asset) Hash() error {
	if len(a.Uris) == 0 {
		return fmt.Errorf("There are no URLs to hash on this Asset.")
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
				return fmt.Errorf("Unable to get asset at %s for hashing.", u)
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
			return fmt.Errorf("Couldn't read from asset at %s, to get its hash.", u)
		}
		// Hash and check against the previous hash.
		hashMap[u] = hex.EncodeToString(s.Sum(nil))
		if i > 0 && hashMap[u] != prevHash {
			return fmt.Errorf("Hash of asset at %s, is not identical to one of the others: %s.", u, prevHash)
		}
		prevHash = hashMap[u]
	}
	// Set hash, having successfully hashing each URL, and checking for duplicates.
	a.Hash = hashMap[a.Uris[0]]

	return nil
}
