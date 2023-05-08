/* Package nft implements basic types and methods for working with NFT assets. */
package nft

// NFT represents an NFT.
type Nft struct {
	*Asset
	Metadata *Asset
	License  *Asset
	Fee      uint    // Fee in mojos
	Royalty  float64 // Royalty as a percentage; not a fraction or basis points.
}
