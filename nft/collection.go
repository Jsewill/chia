package nft

// Collection holds NFT collection data.
type Collection struct {
	Id
	Nfts    []*NFT
	Fee     uint
	Royalty float64
}
