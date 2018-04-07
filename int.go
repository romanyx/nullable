package nullable

import (
	"encoding/json"
	"reflect"
)

// Int represents an int that may be null or not
// present in json at all.
type Int struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid int64
	Value   int64
}

// UnmarshalJSON implements json.Marshaler interface.
func (i *Int) UnmarshalJSON(data []byte) error {
	i.Present = true

	if reflect.DeepEqual(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &i.Value); err != nil {
		return err
	}

	i.Valid = true
	return nil
}

// IntSlice represents an int slice that may be null or not
// present in json at all.
type IntSlice struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid []int64
	Value   []int64
}

// UnmarshalJSON implements json.Marshaler interface.
func (i *IntSlice) UnmarshalJSON(data []byte) error {
	i.Present = true

	if reflect.DeepEqual(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &i.Value); err != nil {
		return err
	}

	i.Valid = true
	return nil
}
