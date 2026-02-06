package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("returns API key when header is valid", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey abc123")

		got, err := GetAPIKey(headers)

		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}

		want := "abc123"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})
	t.Run("returns error when authroization header is missing", func(t *testing.T) {
		headers := http.Header{}
		
		_, err := GetAPIKey(headers)
		
		if err == nil {
			t.Errorf("expected nil error, got %v", err)
		}

		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
		}
	})

}

