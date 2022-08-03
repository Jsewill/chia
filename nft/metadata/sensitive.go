package metadata

import (
	"encoding/json"
	"fmt"
)

// SensitiveContent is used to indicate whether the NFT content is "sensitive". It allows for specifying a boolean value, or a list of topics in the form of a string slice.
// Implements json.Marshaler and json.Unmarshaler interfaces, as special handling is necessary.
type SensitiveContent struct {
	Flag         bool `json:"sensitive_content"`
	ContentTypes []string
}

// Implement json.Marshaler
func (sc *SensitiveContent) MarshalJSON() ([]byte, error) {
	v := make([]byte, 0)
	var err error
	if len(sc.ContentTypes) > 0 {
		// Marshal content types if we have any.
		v, err = json.Marshal(sc.ContentTypes)
	} else {
		// Marshal boolean if we don't.
		v, err = json.Marshal(sc.Flag)
	}

	if err != nil {
		return nil, err
	}
	return v, nil
}

// Implement json.Unmarshaler
func (sc *SensitiveContent) UnmarshalJSON(d []byte) error {
	var v interface{}
	if err := json.Unmarshal(d, &v); err != nil {
		return err
	}
	switch t := v.(type) {
	case bool:
		// Simple boolean.
		sc.Flag = t
	case []interface{}:
		// Should be a string slice.
		for _, s := range t {
			if cts, ok := s.(string); ok {
				// String passed, add it to our list of content types.
				sc.ContentTypes = append(sc.ContentTypes, cts)
			} else {
				// The slice wasn't entirely made of strings.
				return fmt.Errorf("Invalid type while unmarshaling JSON to a %T. Expected string, got %T.", sc, t)
			}
		}
	default:
		// Failed to meet specification. Return an error.
		return fmt.Errorf("Invalid type while unmarshaling JSON to a %T. Expected boolean or [], got %T", sc, t)
	}

	return nil
}
