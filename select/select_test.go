package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Website Race", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error. get: %v", err)
		}

		assertCorrectMessage(t, got, want)
	})

	t.Run("Website Race: Error if Response > 10 seconds", func(t *testing.T) {
		server := makeDelayedServer(11 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, (20 * time.Millisecond))
		if err == nil {
			t.Error("expected an error ")
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
