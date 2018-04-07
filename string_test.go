package nullable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		buf    *bytes.Buffer
		expect String
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: String{
				Present: true,
			},
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":"string"}`),
			expect: String{
				Present: true,
				Valid:   true,
				Value:   "string",
			},
		},
		{
			name:   "empty",
			buf:    bytes.NewBufferString(`{}`),
			expect: String{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value String `json:"value"`
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

func TestStringSlice_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		buf    *bytes.Buffer
		expect StringSlice
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: StringSlice{
				Present: true,
			},
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":["string", "string2"]}`),
			expect: StringSlice{
				Present: true,
				Valid:   true,
				Value:   []string{"string", "string2"},
			},
		},
		{
			name:   "empty",
			buf:    bytes.NewBufferString(`{}`),
			expect: StringSlice{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value StringSlice `json:"value"`
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
