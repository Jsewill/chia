package metadata

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Attribute represents an NFT-level attribute. It implements the json.Marshaler and json.Unmarshaler interfaces, as special handling is necessary.
type Attribute struct {
	Type  string `json:"trait_type"`
	Value string `json:"value"`
	Min   int    `json:"min_value,omitempty"`
	Max   int    `json:"max_value,omitempty"`
}

// Implement json.Unmarshaler
func (a *Attribute) UnmarshalJSON(d []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(d, &m); err != nil {
		return err
	}

	// What did we get for the type and value?
	t := []string{
		"trait_type",
		"value",
	}
	for _, n := range t {
		switch v := m[n].(type) {
		case float64:
			// If a number, assume an int, as per the specification.
			a.Type = strconv.Itoa(int(v))
		case string:
			a.Type = v
		default:
			// Anything else is outside of the specification.
			return fmt.Errorf("Could not unmarshal %v, %T", v, v)
		}
	}

	// Since min and max value are both optional, we should process them as such.
	min, ok := m["min_value"]
	if ok {
		// We have a value.
		if minf, ok := min.(float64); ok {
			a.Min = int(minf)
		} else {
			// Supplied JSON was not a number.
			return fmt.Errorf(`Invalid type assertion while unmarshaling JSON to an %T: supplied "min_value" (of type %T) for "trait_type", %s, was outside the specification.`, a, min, a.Type)
		}
	}
	max, ok := m["max_value"]
	if ok {
		// We have a value.
		if maxf, ok := max.(float64); ok {
			a.Max = int(maxf)
		} else {
			// Supplied JSON was not a number.
			return fmt.Errorf(`Invalid type assertion while unmarshaling JSON to an %T: supplied "max_value" (of type %T) for "trait_type", %s, was outside the specification.`, a, max, a.Type)
		}
	}

	return nil
}

// Implement json.Marshaler
func (a *Attribute) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	// Marshal Type
	tInt, err := strconv.Atoi(a.Type)
	if err != nil {
		// As string
		m["trait_type"] = a.Type
	} else {
		// As int
		m["trait_type"] = tInt
	}
	// Marshal Value
	vInt, err := strconv.Atoi(a.Value)
	if err != nil {
		// As string
		m["value"] = a.Value
	} else {
		// As int
		m["value"] = vInt
	}

	// Supply the remaining variables.
	// Assume values are purposely set, if MaxValue is not zero.
	if a.Max != 0 {
		m["max_value"] = a.Max
		m["min_value"] = a.Min
	}

	return json.Marshal(m)
}
