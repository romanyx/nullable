package nullable

import (
	"bytes"
	"encoding/json"
)

var null = []byte("null")

// String represents a string that may be null or not
// present in json at all.
type String struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid string
	Value   string
}

// UnmarshalJSON implements json.Marshaler interface.
func (s *String) UnmarshalJSON(data []byte) error {
	s.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &s.Value); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

// StringSlice represents a []string that may be null or not
// present in json at all.
type StringSlice struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null
	Value   []string
}

// UnmarshalJSON implements json.Marshaler interface.
func (s *StringSlice) UnmarshalJSON(data []byte) error {
	s.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &s.Value); err != nil {
		return err
	}

	s.Valid = true
	return nil
}
