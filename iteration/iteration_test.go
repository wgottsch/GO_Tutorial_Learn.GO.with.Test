package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Repeating Test", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})

}
func ExampleAdd() {
	got := Repeat("1", 5)
	fmt.Println(got)
	// Output: 11111
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
