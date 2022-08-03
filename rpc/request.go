package rpc

import (
	"encoding/json"
	"fmt"
)

// Caller is implemented by any type that calls a Chia RPC endpoint.
type Caller interface {
	Call(Procedure, []byte) ([]byte, error)
}

// Sender is implemented by any type which can make a request via an Endpoint.
type Sender interface {
	Send(e *Endpoint) (interface{}, error)
	Procedure() Procedure
}

// Generic response type for use in conjunction with UntypedRequest.
type UntypedResponse map[string]interface{}

func NewUntypedResponse() UntypedResponse {
	return make(UntypedResponse)
}

// Generic request type for use when other types won't do, or you don't want to write your own. Implements Sender.
type UntypedRequest struct {
	Proc Procedure
	Data map[string]interface{}
}

func NewUntypedRequest(p Procedure) *UntypedRequest {
	return &UntypedRequest{Proc: p, Data: make(map[string]interface{})}
}

func (u *UntypedRequest) Procedure() Procedure {
	return u.Proc
}

func (u *UntypedRequest) Send(e *Endpoint) (UntypedResponse, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(u.Data)
	if err != nil {
		return nil, err
	}
	// Make request
	out, err := e.Call(u.Proc, j)
	if err != nil {
		return nil, err
	}
	// Handle response
	ur := make(UntypedResponse)
	err = json.Unmarshal(out, ur)
	if err != nil {
		return nil, err
	}
	return UntypedResponse(ur), nil
}

func (u *UntypedRequest) String() string {
	j, err := json.Marshal(u.Data)
	if err != nil {
		// Log error
		fmt.Println(err)
	}
	return fmt.Sprintf(`%s %q`, u.Proc, j)
}
