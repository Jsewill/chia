# Chia

## Chia is a Go (https://golang.org/) module which contains a set of packages for interacting with the Chia Blockchain (https://chia.net/ and https://github.com/Chia-Network/Chia-Blockchain).
### Neither this project, nor its author are affiliated with Chia Network, or the Chia Blockchain, in any way. The name, "Chia" as used in this repository, is simply to keep things clear and Go-like; to indicate what it, and its sub-packages are for, and to make import paths clear.

This module is a work in progress, and is, as such, subject to changes, even breaking ones. It currently houses chia/rpc and chia/nft. See each of these for their respective descriptions and usage (if applicable). Also, check out another work-in-progress of ours, called Artwork: https://github.com/Jsewill/artwork

If you wish to donate toward the development of this project, please use any of the following addresses:

	XCH (Chia): xch1d80tfje65xy97fpxg7kl89wugnd6svlv5uag2qays0um5ay5sn0qz8vph8
	BTC (Bitcoin): 39AcR4aQtvBScT2DEgdqefiH3S8CQzMGfV
	ETH (Ethereum): 0xeA01FC83cee4B89DbD1d27CA0A32bB7B4b1253d1
	USDT (Tether): 0x6be549359f580BC1d8Db61C0CF13198eF14eD999
	USDC (USD Coin): 0xfb9a45a7Fe781D4fc449e63532424bCfDfB086B4

## chia/rpc
### Example Usage

#### Sync Status
```go
r := &rpc.SyncStatusRequest{}
status, err := r.Send(rpc.Wallet)
if err != nil {
    // Sync Status Request failed; handle error.
}
if !status.Success {
    // Sync Status Request was unsuccessful; handle case.
}
if !status.Synced {
    // Check to see if wallet is actively synchronizing.
    if !status.Syncing {
        // Wallet is not actively synchronizing.
    }
}
```

#### Wallet Balance
```go
r := &rpc.WalletBalanceRequest{WalletId: 1}
balance, err := r.Send(rpc.Wallet)
if err != nil {
    // Wallet Balance Request failed; handle error
}
if !balance.Success {
    // Sync Status Request was unsuccessful; handle error.
}
// Got wallet balance.
fmt.Printf("Wallet Balance: %v", balance.WalletBalance)

```
