# RPC

## RPC is a go package meant for interacting with the Chia Blockchain RPC API.

This package is currently a work in progress, and does not yet support much more of the API than is necessary to mint an NFT. This will improve as time allows!

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
