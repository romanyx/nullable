package nullable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestFloat_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		buf    *bytes.Buffer
		expect Float
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: Float{
				Present: true,
			},
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":1.1}`),
			expect: Float{
				Present: true,
				Valid:   true,
				Value:   1.1,
			},
		},
		{
			name:   "empty",
			buf:    bytes.NewBufferString(`{}`),
			expect: Float{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value Float `json:"value"`
			}{}

			if err := json.Unmarshal(tt.buf.Bytes(), &str); err != nil {
				t.Fatalf("unexpected unmarshaling error: %s", err)
			}

			got := str.Value
			if got.Present != tt.expect.Present || got.Valid != tt.expect.Valid || got.Value != tt.expect.Value {
				t.Errorf("expected value to be %#v got %#v", tt.expect, got)
			}
		})
	}
}

func TestFloatSlice_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		buf    *bytes.Buffer
		expect FloatSlice
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: FloatSlice{
				Present: true,
			},
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":[1.1, 1.2]}`),
			expect: FloatSlice{
				Present: true,
				Valid:   true,
				Value:   []float64{1.1, 1.2},
			},
		},
		{
			name:   "empty",
			buf:    bytes.NewBufferString(`{}`),
			expect: FloatSlice{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value FloatSlice `json:"value"`
			}{}

			if err := json.Unmarshal(tt.buf.Bytes(), &str); err != nil {
				t.Fatalf("unexpected unmarshaling error: %s", err)
			}

			got := str.Value
			if got.Present != tt.expect.Present || got.Valid != tt.expect.Valid || fmt.Sprintf("%v", got.Value) != fmt.Sprintf("%v", tt.expect.Value) {
				t.Errorf("expected value to be %#v got %#v", tt.expect, got)
			}
		})
	}
}
