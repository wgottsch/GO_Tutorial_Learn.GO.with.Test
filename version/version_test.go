package version

import "testing"

func TestStructs(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Test Area", func(t *testing.T) {
		got := version()
		want := "00.01.02"
		assertCorrectMessage(t, got, want)
	})
}
