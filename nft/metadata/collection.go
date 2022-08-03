package metadata

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Collection represents a collection and its attributes.
type Collection struct {
	Id         string                 `json:"id"`
	Name       string                 `json:"name"`
	Attributes []*CollectionAttribute `json:"attributes"`
}

// Attribute retrives an attribute by type, from the Attribute slice. Returns an empty *Attribute if not found.
func (ca *Collection) Attribute(t string) *CollectionAttribute {
	for _, a := range ca.Attributes {
		if a.Type == t {
			return a
		}
	}
	return &CollectionAttribute{}
}

// CollectionAttribute represents an NFT collection-level attribute. It implements the json.Marshaler and json.Unmarshaler interfaces, since special cases are needed.
type CollectionAttribute struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Implement json.Unmarshaler
func (a *CollectionAttribute) UnmarshalJSON(d []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(d, &m); err != nil {
		return err
	}

	// What did we get for the type? But for clarity's sake, this could be consolidated into a loop, since both Type and Value require the same treatment.
	switch v := m["type"].(type) {
	case float64:
		// If a number, assume an int, as per the specification.
		a.Type = strconv.Itoa(int(v))
	case string:
		a.Type = v
	default:
		// Anything else is outside of the specification.
		return fmt.Errorf("Could not unmarshal %v", v)
	}

	// What did we get for the value?
	switch v := m["value"].(type) {
	case float64:
		// If a number, assume an int, as per the specification.
		a.Type = strconv.Itoa(int(v))
	case string:
		a.Value = v
	default:
		// Anything else is outside of the specification.
		return fmt.Errorf("Could not unmarshal %v", v)
	}

	return nil
}

// Implement json.Marshaler
func (ca *CollectionAttribute) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	// Marshal Type
	tInt, err := strconv.Atoi(ca.Type)
	if err != nil {
		// As string
		m["type"] = ca.Type
	} else {
		// As int
		m["type"] = tInt
	}
	// Marshal Value
	vInt, err := strconv.Atoi(ca.Value)
	if err != nil {
		// As string
		m["value"] = ca.Value
	} else {
		// As int
		m["value"] = vInt
	}

	return json.Marshal(m)
}
