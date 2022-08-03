package metadata

import (
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed chip-0007_example.json
var example []byte

func TestUnmarshalIntoMetadata(t *testing.T) {
	um := new(Metadata)
	err := json.Unmarshal(example, um)
	if err != nil {
		t.Errorf("Unmarshal of example data into *Metadata, failed: %s", err)
	}
}

func TextMarshalIntoMetadata(t *testing.T) {
	um := new(Metadata)
	err := json.Unmarshal(example, um)
	if err != nil {
		t.Errorf("Unmarshal of example data into *Metadata, failed: %s", err)
	}

	_, err = json.Marshal(um)
	if err != nil {
		t.Errorf("Marshal of *Metadata into JSON, failed: %s", err)
	}

}
