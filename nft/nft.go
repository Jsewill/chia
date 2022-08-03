/* Package nft implements basic types and methods for working with NFT assets. */
package nft

import "fmt"

// NFT represents an NFT.
type NFT struct {
	*Asset
	Metadata *Asset
	License  *Asset
}

// Retrieve and compare hashes for Asset, Metadata, and License files from their respective URLs.
func (nft *NFT) Hash() error {
	// Try hashing Asset files from URLs.
	err := nft.Hash()
	if err != nil {
		return fmt.Errorf("Error while attempting to hash assets: %s", err)
	}
	// Try hashing Metadata files from URLs.
	err := nft.Metadata.Hash()
	if err != nil {
		return fmt.Errorf("Error while attempting to hash metadata assets: %s", err)
	}
	// Try hashing License files from URLs.
	err := nft.License.Hash()
	if err != nil {
		return fmt.Errorf("Error while attempting to hash license assets: %s", err)
	}

	return nil
}
