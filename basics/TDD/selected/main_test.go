package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	got := Racer(slowUrl, fastUrl)
	want := fastUrl
	if got != want {
		t.Errorf("Got: %q Want: %q", got, want)
	}
}
