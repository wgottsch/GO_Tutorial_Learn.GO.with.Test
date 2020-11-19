package integers

import "testing"

func TestAdder(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("adding 2 numbers (2+2)", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertCorrectMessage(t, got, want)
	})

}
