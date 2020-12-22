package dependencyinjection

import (
	"bytes"
	"net/http"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Print Buffer", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Wolle")

		got := buffer.String()
		want := "Hello, Wolle"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Print HTTP Response", func(t *testing.T) {
		http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	})

	t.Run("print Stout", func(t *testing.T) {
		Greet(os.Stdout, "Lizzy")
	})

}
