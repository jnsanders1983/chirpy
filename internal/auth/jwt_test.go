package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWT(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"

	tests := []struct {
		name           string
		expiresIn      time.Duration
		validateSecret string
		wantError      bool
	}{
		{
			name:           "valid token",
			expiresIn:      time.Hour,
			validateSecret: secret,
		},
		{
			name:           "expired token",
			expiresIn:      -time.Hour,
			validateSecret: secret,
			wantError:      true,
		},
		{
			name:           "wrong secret",
			expiresIn:      time.Hour,
			validateSecret: "wrong-secret",
			wantError:      true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			token, err := MakeJWT(userID, secret, tc.expiresIn)
			if err != nil {
				t.Fatalf("MakeJWT failed: %v", err)
			}

			gotID, err := ValidateJWT(token, tc.validateSecret)
			if tc.wantError {
				if err == nil {
					t.Fatal("expected validation to fail")
				}
				return
			}

			if err != nil {
				t.Fatalf("ValidateJWT failed: %v", err)
			}

			if gotID != userID {
				t.Errorf("got user ID %v, want %v", gotID, userID)
			}
		})
	}
}
