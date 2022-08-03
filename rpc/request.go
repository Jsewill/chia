package rpc

import (
	"encoding/json"
	"fmt"
)

// Procedure represents procedure, and its value is its name.
type Procedure string

// String implements the fmt.Stringer interface.
func (p Procedure) String() string {
	return string(p)
}

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

// Convenience function. May be deprecated.
func NewUntypedResponse() UntypedResponse {
	return make(UntypedResponse)
}

// Generic request type for use when other types won't do, or you don't want to write your own. Implements Sender.
type UntypedRequest struct {
	Proc Procedure
	Data map[string]interface{}
}

// NewUntypedRequest takes a Procedure and returns an UntypedRequest.
func NewUntypedRequest(p Procedure) *UntypedRequest {
	return &UntypedRequest{Proc: p, Data: make(map[string]interface{})}
}

// Procedure returns the procedure currently assigned to this generic request.
func (u *UntypedRequest) Procedure() Procedure {
	return u.Proc
}

// Send sends the request via an Endpoint, and returns the result as an UntypedResponse, or nil, and an error. If successful, error returns nil.
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

// String implements the fmt.Stringer interface.
func (u *UntypedRequest) String() string {
	j, err := json.Marshal(u.Data)
	if err != nil {
		// Log error
		fmt.Println(err)
	}
	return fmt.Sprintf(`%s %q`, u.Proc, j)
}
