package nft

// Collection holds NFT collection data.
type Collection struct {
	Id
	Nfts    []*NFT
	Fee     uint    // Fee in mojos
	Royalty float64 // Royalty as a percentage; not a fraction or basis points.
}
