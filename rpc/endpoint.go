package rpc

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

// NewEndpoint returns and initializes a new *Endpoint. It returns an error on initialization faliure.
func NewEndpoint(n string, h string, p uint) (*Endpoint, error) {
	e := &Endpoint{
		Name: n,
		Host: h,
		Port: p,
	}
	err := e.Init()
	return e, err
}

// An Endpoint represents a Chia RPC endpoint. It implements Caller.
type Endpoint struct {
	Name string
	Host string
	Port uint
	*http.Transport
	*http.Client
}

// Initializes the Endpoint's HTTP Transport and Client properties.
func (e *Endpoint) Init() error {
	// Compile SSL file paths from default/custom data.
	rp := filepath.Join(DefaultPath, "mainnet/config/ssl/ca", "private_ca.crt")
	cp := filepath.Join(DefaultPath, "mainnet/config/ssl", e.Name, "private_"+e.Name+".crt")
	kp := filepath.Join(DefaultPath, "mainnet/config/ssl", e.Name, "private_"+e.Name+".key")

	// Load Certs.
	c, err := tls.LoadX509KeyPair(cp, kp)
	if err != nil {
		return err
	}

	// Get Chia CA.
	ownCa, err := ioutil.ReadFile(rp)
	if err != nil {
		logErr.Println(err)
	}

	// Make pool from Chia CA.
	caRoots := x509.NewCertPool()
	caRoots.AppendCertsFromPEM(ownCa)

	// Setup Transport with TLS.
	e.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{c},
			// @TODO: Get custom (chia) Cert pool working properly so we don't need this.
			InsecureSkipVerify: true,
			RootCAs:            caRoots,
		},
	}
	// Setup Client.
	e.Client = &http.Client{Transport: e.Transport}
	return nil
}

// Post wraps the embedded *http.Client method. Used by Endpoint.Call() to make a POST request to a Chia RPC endpoint.
func (e *Endpoint) Post(p Procedure, b io.Reader) (*http.Response, error) {
	uri := strings.Join([]string{e.String(), string(p)}, "/")
	r, err := e.Client.Post(uri, "application/json", b)
	logErr.Println(err)
	return r, err
}

// Returns the Endpoint URI. Implements the fmt.Stringer interface.
func (e *Endpoint) String() string {
	return fmt.Sprintf("https://%s:%d", e.Host, e.Port)
}

// Call makes a call to the Chia RPC endpoint and returns the response as a byte slice.
// Takes a Procedure, p, and a payload as a JSON byte slice, j.
func (e *Endpoint) Call(p Procedure, j []byte) ([]byte, error) {
	// Make POST request
	buf := bytes.NewReader(j)
	r, err := e.Post(p, buf)
	if err != nil {
		err = fmt.Errorf("Error with POST request to %s : %v", e, err)
		logErr.Println(err)
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
