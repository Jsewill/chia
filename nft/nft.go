/* Package nft implements basic types and methods for working with NFT assets. */
package nft

// NFT represents an NFT.
type Nft struct {
	*Asset
	Metadata *Asset
	License  *Asset
	Fee      uint
	Royalty  float64
}
