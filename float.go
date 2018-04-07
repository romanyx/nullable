package nullable

import (
	"encoding/json"
	"reflect"
)

// Float represents a float that may be null or not
// present in json at all.
type Float struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid float
	Value   float64
}

// UnmarshalJSON implements json.Marshaler interface.
func (f *Float) UnmarshalJSON(data []byte) error {
	f.Present = true

	if reflect.DeepEqual(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &f.Value); err != nil {
		return err
	}

	f.Valid = true
	return nil
}

// FloatSlice represents a float slice that may be null or not
// present in json at all.
type FloatSlice struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid []float64
	Value   []float64
}

// UnmarshalJSON implements json.Marshaler interface.
func (f *FloatSlice) UnmarshalJSON(data []byte) error {
	f.Present = true

	if reflect.DeepEqual(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &f.Value); err != nil {
		return err
	}

	f.Valid = true
	return nil
}
