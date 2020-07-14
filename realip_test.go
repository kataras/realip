package realip

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	original := "80.106.234.17"
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("X-Forwarded-For", "80.106.234.17")

	if expected, got := original, Get(r); expected != got {
		t.Fatalf("expected: %s but got: %s", expected, got)
	}

	// range.End is valid for public address.
	expectedEndIsValid := "10.255.255.255"
	r.Header.Set("X-Real-Ip", expectedEndIsValid)
	if expected, got := expectedEndIsValid, Get(r); expected != got {
		t.Fatalf("expected: %s but got: %s", expected, got)
	}

	// range.Start >= value < End is private.
	beforeEndIsInvalid := "10.255.255.254"
	r.Header.Set("X-Real-Ip", beforeEndIsInvalid)
	if expected, got := original, Get(r); expected != got {
		t.Fatalf("expected: %s but got: %s", expected, got)
	}
}
