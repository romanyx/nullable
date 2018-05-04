package nullable

import (
	"bytes"
	"encoding/json"
)

// Bool represents a bool that may be null or not
// present in json at all.
type Bool struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid bool
	Value   bool
}

// UnmarshalJSON implements json.Marshaler interface.
func (b *Bool) UnmarshalJSON(data []byte) error {
	b.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &b.Value); err != nil {
		return err
	}

	b.Valid = true
	return nil
}
