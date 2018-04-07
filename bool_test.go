package nullable

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestBool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		buf    *bytes.Buffer
		expect Bool
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: Bool{
				Present: true,
			},
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":true}`),
			expect: Bool{
				Present: true,
				Valid:   true,
				Value:   true,
			},
		},
		{
			name:   "empty",
			buf:    bytes.NewBufferString(`{}`),
			expect: Bool{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value Bool `json:"value"`
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
