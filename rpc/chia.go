/* Package rpc provides types and methods for use in communicating with the Chia RPC */
package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	defaultHost = "localhost"
)

var (
	DefaultPath     = ".chia"
	DefaultCertPath = "mainnet/config/ssl/"
	HomeDir         = ""
)

func init() {
	HomeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Couldn't get home directory path. Error: " + fmt.Sprint(err))
	}
	if _, err := os.Stat(HomeDir); os.IsNotExist(err) {
		panic("System specified home directory does not exist.")
	}
	DefaultPath = filepath.Join(HomeDir, DefaultPath)
}

// Call is a general endpoint call function for convenience.
// @TODO: Expand definition for use with interfaces. May require some renaming of Endpoint and Procedure functions as well as package interfaces rework.
func Call(e Endpoint, p Procedure, d interface{}) ([]byte, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	// Make POST request
	buf := bytes.NewReader(j)
	r, err := e.Post(p, buf)
	if err != nil {
		err = fmt.Errorf("Error with POST request to %s : %v", e, err)
		return nil, err
	}
	defer r.Body.Close()
	// Read response
	b, err := io.ReadAll(r.Body)
	// Return if error or status code shows an unsuccessful request.
	if err != nil || r.StatusCode > 299 {
		return b, err
	}

	return b, nil
}
