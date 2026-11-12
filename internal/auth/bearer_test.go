package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name      string
		header    string
		want      string
		wantError bool
	}{
		{"valid", "Bearer abc123", "abc123", false},
		{"extra whitespace", "  Bearer   abc123  ", "abc123", false},
		{"missing header", "", "", true},
		{"missing token", "Bearer", "", true},
		{"wrong scheme", "Basic abc123", "", true},
		{"extra token", "Bearer abc123 xyz", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headers := make(http.Header)
			if tc.header != "" {
				headers.Set("Authorization", tc.header)
			}

			got, err := GetBearerToken(headers)
			if (err != nil) != tc.wantError {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
