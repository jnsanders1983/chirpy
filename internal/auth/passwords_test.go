package auth

import "testing"

func TestPasswordHash(t *testing.T) {
	hash, err := HashPassword("04234")
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name     string
		password string
		want     bool
	}{
		{"correct password", "04234", true},
		{"incorrect password", "wrong", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			match, err := CheckPasswordHash(tc.password, hash)
			if err != nil {
				t.Fatal(err)
			}
			if match != tc.want {
				t.Errorf("got %v, want %v", match, tc.want)
			}
		})
	}
}

func TestInvalidHash(t *testing.T) {
	match, err := CheckPasswordHash("04234", "unset")
	if err == nil || match {
		t.Fatal("expected an error and no match for an invalid hash")
	}
}
