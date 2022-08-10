package rpc

import (
	"encoding/json"
	"testing"
)

func TestCallSyncStatus(t *testing.T) {
	// Make request
	r := &SyncStatusRequest{}
	out, err := Call(Wallet, r.Procedure().String(), r)
	if err != nil {
		t.Errorf("Sync Status Request failed: %s", err)
	}
	// Handle response
	status := new(SyncStatusResponse)
	err = json.Unmarshal(out, status)
	if err != nil {
		t.Errorf("Sync Status unmarshal response failed: %s", err)
	}
	if !status.Success {
		t.Errorf("Sync Status Request was unsuccessful: %s", status.Error)
	}
	t.Log(status, err)
}

func TestCallGenericSyncStatus(t *testing.T) {
	// Make request
	r := NewUntypedRequest(WalletSyncStatus)
	out, err := Call(Wallet, r.Procedure().String(), &r)
	if err != nil {
		t.Errorf("Sync Status Request failed: %s", err)
	}
	// Handle response
	status := NewUntypedResponse()
	err = json.Unmarshal(out, &status)
	if err != nil {
		t.Errorf("Sync Status unmarshal response failed: %s", err)
	}
	if success, ok := status["success"].(bool); ok && !success {
		t.Errorf("Sync Status Request was unsuccessful: %s", status["error"])
	} else if !ok {
		t.Errorf("Sync Status Response success not boolean: %s; error: %s", status["success"], status["error"])

	}
	t.Log(status, err)
}
