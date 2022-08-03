package rpc

import (
	"encoding/json"
	"fmt"
)

type Procedure string

const (
	WalletMintNFT    Procedure = "nft_mint_nft"
	WalletSyncStatus Procedure = "get_sync_status"
	WalletGetBalance Procedure = `get_wallet_balance`
)

var (
	Wallet *Endpoint = &Endpoint{Name: "wallet", Host: defaultHost, Port: 9256}
)

func init() {
	err := Wallet.Init()
	if err != nil {
		panic(err)
	}
}

type MintResponse struct {
	Spend_bundle *SpendBundle
	Success      bool
	Wallet_id    uint
}

type MintRequest struct {
	WalletId          int      `json:"wallet_id"`
	Uris              []string `json:"uris"`
	Hash              string   `json:"hash"`
	DidId             string   `json:"did_id,omitempty"`
	MetaUris          []string `json:"meta_uris,omitempty"`
	MetaHash          string   `json:"meta_hash,omitempty"`
	LicenseUris       []string `json:"license_uris,omitempty"`
	LicenseHash       string   `json:"license_hash,omitempty"`
	RoyaltyAddress    string   `json:"royalty_address,omitempty"`
	RoyaltyPercentage int      `json:"royalty_percentage,omitempty"`
	TargetAddress     string   `json:"target_address,omitempty"`
	Fee               float64  `json:"fee,omitempty"`
	SeriesNumber      int      `json:"series_number,omitempty"`  // In CHIP-0007, but not in chia 1.4.0 as "series_number", due to bug. This has been deprecated as of chia 1.5.0, likely until NFT2.
	SeriesTotal       int      `json:"series_total,omitempty"`   // In CHIP-0007, but not in chia 1.4.0 as "series_total", due to bug.. This has been deprecated as of chia 1.5.0, likely until NFT2.
	EditionNumber     int      `json:"edition_number,omitempty"` // Not in CHIP-0007, but in chia 1.4.0 as "series_total", rather than "edition_total", due to bug. This was fixed in chia 1.5.0.
	EditionTotal      int      `json:"edition_total,omitempty"`  // Not in CHIP-0007, but in chia 1.4.0 as "series_number", rather than "edition_total", due to bug. This was fixed in chia 1.5.0.
}

func (m *MintRequest) Procedure() Procedure {
	return WalletMintNFT
}

func (m MintRequest) Send(e *Endpoint) (*MintResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	// Make request
	out, err := e.Call(m.Procedure(), j)
	if err != nil {
		return nil, err
	}
	// Handle response
	mr := new(MintResponse)
	err = json.Unmarshal(out, mr)
	if err != nil {
		return nil, err
	}
	return mr, nil
}

func (m *MintRequest) String() string {
	j, err := json.Marshal(m)
	if err != nil {
		// Log error
		fmt.Println(err)
	}
	return fmt.Sprintf(`%s %q`, m.Procedure(), j)
}

type Solution struct {
	Coin         *Coin  `json:"coin"`
	PuzzleReveal string `json:"puzzle_reveal"`
	Solution     string `json:"solution"`
}

type SpendBundle struct {
	AggregatedSignature string      `json:"aggregated_signature"`
	CoinSolutions       []*Solution `json:"coin_solutions"`
}

type SyncStatusResponse struct {
	GenesisInitialized bool   `json:"genesis_initialized"`
	Success            bool   `json:"success"`
	Synced             bool   `json:"synced"`
	Syncing            bool   `json:"syncing"`
	Error              string `json:"error"`
}

type SyncStatusRequest struct{}

func (s *SyncStatusRequest) Procedure() Procedure {
	return WalletSyncStatus
}

func (s *SyncStatusRequest) Send(e *Endpoint) (*SyncStatusResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	// Make request
	out, err := e.Call(s.Procedure(), j)
	if err != nil {
		return nil, err
	}
	// Handle response
	sr := new(SyncStatusResponse)
	err = json.Unmarshal(out, sr)
	if err != nil {
		return nil, err
	}
	return sr, nil
}

func (s *SyncStatusRequest) String() string {
	j, err := json.Marshal(s)
	if err != nil {
		// Log error
		fmt.Println(err)
	}
	return fmt.Sprintf(`%s %q`, s.Procedure(), j)
}

type WalletBalanceRequest struct {
	Wallet_id uint `json:"wallet_id"`
}

func (w *WalletBalanceRequest) Procedure() Procedure {
	return WalletGetBalance
}

func (w *WalletBalanceRequest) Send(e *Endpoint) (*WalletBalanceResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(w)
	if err != nil {
		return nil, err
	}
	// Make request
	out, err := e.Call(w.Procedure(), j)
	if err != nil {
		return nil, err
	}
	// Handle response
	wr := new(WalletBalanceResponse)
	err = json.Unmarshal(out, wr)
	if err != nil {
		return nil, err
	}
	return wr, nil
}

func (w *WalletBalanceRequest) String() string {
	j, err := json.Marshal(w)
	if err != nil {
		// Log error
		fmt.Println(err)
	}
	return fmt.Sprintf(`%s %q`, w.Procedure(), j)
}

type WalletBalance struct {
	ConfirmedWalletBalance   uint `json:"confirmed_wallet_balance"`
	Fingerprint              uint `json:"fingerprint"`
	MaxSendAmount            uint `json:"max_send_amount"`
	PendingChange            uint `json:"pending_change"`
	PendingCoinRemovalCount  uint `json:"pending_coin_removal_count"`
	SpendableBalance         uint `json:"spendable_balance"`
	UnconfirmedWalletBalance uint `json:"unconfirmed_wallet_balance"`
	UnspentCoinCount         uint `json:"unspent_coin_amount"`
	WalletId                 uint `json:"wallet_id"`
}

type WalletBalanceResponse struct {
	WalletBalance *WalletBalance `json:"wallet_balance"`
	Success       bool           `json:"success"`
	Error         string         `json:"error"`
}
