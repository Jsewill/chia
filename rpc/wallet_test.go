package rpc

import "testing"

func TestSyncStatus(t *testing.T) {
	r := &SyncStatusRequest{}
	status, err := r.Send(Wallet)
	if err != nil {
		t.Errorf("Sync Status Request failed: %s", err)
	}
	if !status.Success {
		t.Errorf("Sync Status Request was unsuccessful: %s", status.Error)
	}
	t.Log(status, err)
}

func TestWalletBalance(t *testing.T) {
	r := &WalletBalanceRequest{WalletId: 1}
	balance, err := r.Send(Wallet)
	if err != nil {
		t.Errorf("Wallet Balance Request failed: %s", err)
	}
	if !balance.Success {
		t.Errorf("Sync Status Request was unsuccessful: %s", balance.Error)
	}
	t.Log(balance, err)
}
