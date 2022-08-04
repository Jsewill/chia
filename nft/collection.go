package nft

type Collection struct {
	ID
	NFTs    []*NFT
	Fee     uint
	Royalty float64
}
