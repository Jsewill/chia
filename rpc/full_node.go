package rpc

import (
	"encoding/json"
	"fmt"
)

const (
	FullNodeCoinRecordByName      Procedure = "get_coin_record_by_name"
	FullNodeCoinRecordByNames     Procedure = "get_coin_record_by_names"
	FullNodeCoinRecordByParentIds Procedure = "get_coin_record_by_parent_ids"
	FullNodeCoinRecordByHints     Procedure = "get_coin_record_by_hints"
)

var (
	FullNode *Endpoint = &Endpoint{Name: "full_node", Host: defaultHost, Port: 8555}
)

func init() {
	err := FullNode.Init()
	if err != nil {
		logErr.Panicln(err)
	}
}

/*
get_network_info
get_blockchain_state
get_block
get_blocks
get_block_count_metrics
get_block_record_by_height
get_block_record
get_block_records
get_block_spends
get_unfinished_block_headers
get_network_space
get_additions_and_removals
get_puzzle_and_solution
get_recent_signage_point_or_eos
get_coin_records_by_puzzle_hash
get_coin_records_by_puzzle_hashes
push_tx
get_all_mempool_tx_ids
get_all_mempool_items
get_mempool_item_by_tx_id
get_routes
healthz
*/

// Coin contains details about a specific coin.
type Coin struct {
	Amount         uint   `json:"amount"`
	ParentCoinInfo string `json:"parent_coin_info"`
	PuzzleHash     string `json:"puzzle_hash"`
}

// CoinRecord contains details about a coin record.
type CoinRecord struct {
	Coin                *Coin `json:"coin"`
	Coinbase            bool  `json:"coinbase"`
	ConfirmedBlockIndex uint  `json:"confirmed_block_index"`
	Spent               bool  `json:"spent"`
	SpentBlockIndex     uint  `json:"spent_block_index"`
	Timestamp           uint  `json:"timestamp"`
}

// CoinRecordResponse represents the Chia RPC API's response to a CoinRecordRequest.
type CoinRecordResponse struct {
	CoinRecord *CoinRecord `json:"coin_record"`
	Success    bool        `json:"success"`
	Error      string      `json:"error"`
}

// CoinRecordRequest is a type for making a request for a single CoinRecord by name.
type CoinRecordRequest struct {
	Name string `json:"name"`
}

// Procedure returns the Procedure which this request will use.
func (c *CoinRecordRequest) Procedure() Procedure {
	return FullNodeCoinRecordByName
}

// Sends the request via an Endpoint, and returns the response, and an error. If successful, error returns nil.
func (c *CoinRecordRequest) Send(e *Endpoint) (*CoinRecordResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(c)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Make request
	out, err := e.Call(c.Procedure(), j)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Handle response
	cr := new(CoinRecordResponse)
	err = json.Unmarshal(out, cr)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	return cr, nil
}

// String implements the fmt.Stringer interface.
func (c *CoinRecordRequest) String() string {
	j, err := json.Marshal(c)
	if err != nil {
		logErr.Println(err)
	}
	return fmt.Sprintf(`%s %q`, c.Procedure(), j)
}

// CoinRecordsResponse represents the Chia RPC API's response to a CoinRecordRequest.
type CoinRecordsResponse struct {
	CoinRecords []*CoinRecord `json:"coin_records"`
	Success     bool          `json:"success"`
	Error       string        `json:"error"`
}

// CoinRecordsRequest is a type for making at least one request for a multiple CoinRecords by coin names, parent ids, and/or hints.
type CoinRecordsRequest struct {
	Names        []string `json:"names"`
	ParentIds    []string `json:"parent_ids"`
	Hints        []string `json:"hints"`
	StartHeight  uint     `json:"start_height,omitempty"`
	EndHeight    uint     `json:"end_height,omitempty"`
	IncludeSpent bool     `json:"include_spent_coins,omitempty"`
}

// Procedure returns the Procedure which this request will use. Though this type is designed to call multiple procedures, this method will always return FullNodeCoinRecordByNames.
func (c *CoinRecordsRequest) Procedure() Procedure {
	return FullNodeCoinRecordByNames
}

// Sends at least one request, of at least one CoinRecordsByN procedure, via an Endpoint, and returns the response, and any error. If successful, error returns nil.
func (c *CoinRecordsRequest) Send(e *Endpoint) (*CoinRecordsResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(c)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Make request(s)
	responses := make([][]byte, 0)
	switch {
	case len(c.Names) > 0:
		// Make request with Names
		out, err := e.Call(c.Procedure(), j)
		if err != nil {
			logErr.Println(err)
			return nil, err
		}
		responses = append(responses, out)
		fallthrough
	case len(c.ParentIds) > 0:
		// Make request with ParentIds
		out, err := e.Call(FullNodeCoinRecordByParentIds, j)
		if err != nil {
			logErr.Println(err)
			return nil, err
		}
		responses = append(responses, out)
		fallthrough
	case len(c.Hints) > 0:
		// Make request with Hints
		out, err := e.Call(FullNodeCoinRecordByHints, j)
		if err != nil {
			logErr.Println(err)
			return nil, err
		}
		responses = append(responses, out)
	default:
		// Nothing to request.
		err = fmt.Errorf("Failed to make CoinRecords request, please set Names, ParentIds, or Hints.")
		logErr.Println(err)

		return nil, err
	}
	// Handle can consolidate response(s)
	errs := NewErrors()
	cr := new(CoinRecordsResponse)
	for _, rout := range responses {
		// Handle response
		tempCr := new(CoinRecordsResponse)
		err = json.Unmarshal(rout, tempCr)
		if err != nil {
			errs = append(errs, fmt.Errorf("CoinRecordsRequest failed to unmarshal response. Error: %s", err))
			continue
		}
		if !tempCr.Success {
			errs = append(errs, fmt.Errorf("CoinRecordsRequest was unsuccessful. Error: %s", tempCr.Error))
			continue
		}
		cr.Success = true
		cr.CoinRecords = append(cr.CoinRecords, tempCr.CoinRecords...)
	}
	if len(errs) > 0 {
		err = fmt.Errorf("%s \n", errs)
		logErr.Println(err)
		return cr, errs
	}

	return cr, err
}

// String implements the fmt.Stringer interface.
func (c *CoinRecordsRequest) String() string {
	j, err := json.Marshal(c)
	if err != nil {
		logErr.Println(err)
	}
	return fmt.Sprintf(`%s %q`, c.Procedure(), j)
}

// CoinRecordsByNameRequest is a type for making a request for a multiple CoinRecords by coin name.
type CoinRecordsByNameRequest struct {
	Name         []string `json:"parent_ids"`
	StartHeight  uint     `json:"start_height,omitempty"`
	EndHeight    uint     `json:"end_height,omitempty"`
	IncludeSpent bool     `json:"include_spent_coins,omitempty"`
}

// Procedure returns the Procedure which this request will use.
func (c *CoinRecordsByNameRequest) Procedure() Procedure {
	return FullNodeCoinRecordByName
}

// Sends the request via an Endpoint, and returns the response, and an error. If successful, error returns nil.
func (c *CoinRecordsByNameRequest) Send(e *Endpoint) (*CoinRecordsResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(c)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Make request
	out, err := e.Call(c.Procedure(), j)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Handle response
	cr := new(CoinRecordsResponse)
	err = json.Unmarshal(out, cr)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	return cr, nil
}

// String implements the fmt.Stringer interface.
func (c *CoinRecordsByNameRequest) String() string {
	j, err := json.Marshal(c)
	if err != nil {
		logErr.Println(err)
	}
	return fmt.Sprintf(`%s %q`, c.Procedure(), j)
}

// CoinRecordsRequest is a type for making a request for a multiple CoinRecords by parent ids.
type CoinRecordsByParentIdsRequest struct {
	ParentIds    []string `json:"parent_ids"`
	StartHeight  uint     `json:"start_height,omitempty"`
	EndHeight    uint     `json:"end_height,omitempty"`
	IncludeSpent bool     `json:"include_spent_coins,omitempty"`
}

// Procedure returns the Procedure which this request will use.
func (c *CoinRecordsByParentIdsRequest) Procedure() Procedure {
	return FullNodeCoinRecordByParentIds
}

// Sends the request via an Endpoint, and returns the response, and an error. If successful, error returns nil.
func (c *CoinRecordsByParentIdsRequest) Send(e *Endpoint) (*CoinRecordsResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(c)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Make request
	out, err := e.Call(c.Procedure(), j)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	// Handle response
	cr := new(CoinRecordsResponse)
	err = json.Unmarshal(out, cr)
	if err != nil {
		logErr.Println(err)
		return nil, err
	}
	return cr, nil
}

// String implements the fmt.Stringer interface.
func (c *CoinRecordsByParentIdsRequest) String() string {
	j, err := json.Marshal(c)
	if err != nil {
		logErr.Println(err)
	}
	return fmt.Sprintf(`%s %q`, c.Procedure(), j)
}

// @TODO: implement push_tx for bulk minting.
