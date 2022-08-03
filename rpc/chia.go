/* Package rpc provides types and methods for use in communicating with the Chia RPC */
package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
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
// It takes a Caller, procedure name string, and an any data type to marshal into JSON.
// @TODO: Expand definition for use with interfaces. May require some renaming of Endpoint and Procedure functions as well as package interfaces rework.
func Call(c Caller, p string, d interface{}) ([]byte, error) {
	// Marshal request body as JSON
	j, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	// Make POST request
	buf := bytes.NewReader(j)
	return c.Call(Procedure(p), buf)
}
