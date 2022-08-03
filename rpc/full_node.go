package rpc

import (
	"encoding/json"
	"fmt"
)

const (
	FullNodeCoinRecordByName Procedure = "get_coin_record_by_name"
)

var (
	FullNode *Endpoint = &Endpoint{Name: "full_node", Host: defaultHost, Port: 8555}
)

func init() {
	err := FullNode.Init()
	if err != nil {
		panic(err)
	}
}

type Coin struct {
	Amount         uint   `json:"amount"`
	ParentCoinInfo string `json:"parent_coin_info"`
	PuzzleHash     string `json:"puzzle_hash"`
}

type CoinRecord struct {
	Coin                *Coin `json:"coin"`
	Coinbase            bool  `json:"coinbase"`
	ConfirmedBlockIndex uint  `json:"confirmed_block_index"`
	Spent               bool  `json:"spent"`
	SpentBlockIndex     uint  `json:"spent_block_index"`
	Timestamp           uint  `json:"timestamp"`
}

type CoinRecordResponse struct {
	CoinRecord *CoinRecord `json:"coin_record"`
	Success    bool        `json:"success"`
	Error      string      `json:"error"`
}

type CoinRecordRequest struct {
	Name string `json:"name"`
}

func (c *CoinRecordRequest) Procedure() Procedure {
	return FullNodeCoinRecordByName
}

func (c *CoinRecordRequest) Send(e *Endpoint) (*CoinRecordResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	// Make request
	out, err := e.Call(c.Procedure(), j)
	if err != nil {
		return nil, err
	}
	// Handle response
	cr := new(CoinRecordResponse)
	err = json.Unmarshal(out, cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (c *CoinRecordRequest) String() string {
	j, err := json.Marshal(c)
	if err != nil {
		// Log error
		fmt.Println(err)
	}
	return fmt.Sprintf(`%s %q`, c.Procedure(), j)
}
