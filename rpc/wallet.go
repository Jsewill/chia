package rpc

import (
	"encoding/json"
	"fmt"
)

const (
	WalletNFTMint         Procedure = "nft_mint_nft"
	WalletNFTMintBulk     Procedure = "nft_mint_bulk"
	WalletNFTGetWalletDID Procedure = "nft_get_wallet_did"
	WalletSyncStatus      Procedure = "get_sync_status"
	WalletGetBalance      Procedure = `get_wallet_balance`
	WalletPushTx          Procedure = `push_tx`
)

var (
	Wallet *Endpoint = &Endpoint{Name: "wallet", Host: defaultHost, Port: 9256}
)

func init() {
	err := Wallet.Init()
	if err != nil {
		logErr.Panicln(err)
	}
}

type MetadataListItem struct {
	Uris          []string `json:"uris"`
	MetaUris      []string `json:"meta_uris,omitempty"`
	LicenseUris   []string `json:"license_uris,omitempty"`
	Hash          string   `json:"hash"`
	MetaHash      string   `json:"meta_hash,omitempty"`
	LicenseHash   string   `json:"license_hash,omitempty"`
	EditionNumber int      `json:"edition_number,omitempty"` // Not in CHIP-0007, but in chia 1.4.0 as "series_total", rather than "edition_total", due to bug. This was fixed in chia 1.5.0.
	EditionTotal  int      `json:"edition_total,omitempty"`  // Not in CHIP-0007, but in chia 1.4.0 as "series_number", rather than "edition_total", due to bug. This was fixed in chia 1.5.0.
}

type MintBulkResponse struct {
	NftIdList   []string     `json:"nft_id_list"`
	SpendBundle *SpendBundle `json:"spend_bundle"`
	Success     bool         `json:"success"`
	Error       string       `json:"error"`
}

type MintBulkRequest struct {
	WalletId            int                    `json:"wallet_id"`
	MetadataList        []*MetadataListItem    `json:"metadata_list"`
	RoyaltyPercentage   int                    `json:"royalty_percentage,omitempty"`
	RoyaltyAddress      string                 `json:"royalty_address,omitempty"`
	TargetAddressList   []string               `json:"target_address_list,omitempty"`
	MintNumberStart     int                    `json:"mint_number_start,omitempty"`
	MintTotal           int                    `json:"mint_total,omitempty"`
	XchCoinList         []string               `json:"xch_coin_list,omitempty"`
	XchChangeTarget     string                 `json:"xch_change_target,omitempty"`
	NewInnerPuzHash     string                 `json:"new_innerpuzhash,omitempty"`
	NewP2PuzHash        string                 `json:"new_p2_puzhash,omitempty"`
	DidCoinDict         map[string]interface{} `json:"did_coin_dict,omitempty"`
	DidLineageParentHex string                 `json:"did_lineage_parent_hex,omitempty"`
	MintFromDid         bool                   `json:"mint_from_did,omitempty"`
	Fee                 float64                `json:"fee,omitempty"`
	ReusePuzHash        bool                   `json:"reuse_puzhash,omitempty"`
}

func (m *MintBulkRequest) Procedure() Procedure {
	return WalletNFTMintBulk
}

func (m MintBulkRequest) Send(e *Endpoint) (*MintBulkResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(m)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Make request
	out, err := e.Call(m.Procedure(), j)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Handle response
	mbr := new(MintBulkResponse)
	err = json.Unmarshal(out, mbr)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	return mbr, nil
}

func (m *MintBulkRequest) String() string {
	j, err := json.Marshal(m)
	if err != nil {
		logErr.Println(err)
	}
	return fmt.Sprintf(`%s %q`, m.Procedure(), j)
}

type MintResponse struct {
	Spend_bundle *SpendBundle
	Success      bool   `json:"success"`
	WalletId     uint   `json:"wallet_id"`
	Error        string `json:"error"`
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
	return WalletNFTMint
}

func (m MintRequest) Send(e *Endpoint) (*MintResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(m)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Make request
	out, err := e.Call(m.Procedure(), j)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Handle response
	mr := new(MintResponse)
	err = json.Unmarshal(out, mr)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	return mr, nil
}

func (m *MintRequest) String() string {
	j, err := json.Marshal(m)
	if err != nil {
		logErr.Println(err)
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
		logErr.Println(err)
		return nil, err
	}
	// Make request
	out, err := e.Call(s.Procedure(), j)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Handle response
	sr := new(SyncStatusResponse)
	err = json.Unmarshal(out, sr)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	return sr, nil
}

func (s *SyncStatusRequest) String() string {
	j, err := json.Marshal(s)
	if err != nil {
		logErr.Println(err)
	}
	return fmt.Sprintf(`%s %q`, s.Procedure(), j)
}

type WalletBalanceRequest struct {
	WalletId uint `json:"wallet_id"`
}

func (w *WalletBalanceRequest) Procedure() Procedure {
	return WalletGetBalance
}

func (w *WalletBalanceRequest) Send(e *Endpoint) (*WalletBalanceResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(w)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Make request
	out, err := e.Call(w.Procedure(), j)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Handle response
	wr := new(WalletBalanceResponse)
	err = json.Unmarshal(out, wr)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	return wr, nil
}

func (w *WalletBalanceRequest) String() string {
	j, err := json.Marshal(w)
	if err != nil {
		logErr.Println(err)
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

type NftWalletGetDidRequest struct {
	WalletId uint `json:"wallet_id"`
}

func (n *NftWalletGetDidRequest) Procedure() Procedure {
	return WalletNFTGetWalletDID
}

func (n *NftWalletGetDidRequest) Send(e *Endpoint) (*NftWalletGetDidResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(n)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Make request
	out, err := e.Call(n.Procedure(), j)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Handle response
	nr := new(NftWalletGetDidResponse)
	err = json.Unmarshal(out, nr)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	return nr, nil
}

func (n *NftWalletGetDidRequest) String() string {
	j, err := json.Marshal(n)
	if err != nil {
		logErr.Println(err)
	}
	return fmt.Sprintf(`%s %q`, n.Procedure(), j)
}

type NftWalletGetDidResponse struct {
	DidId   int    `json:"did_id"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
