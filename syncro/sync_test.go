package syncro

import (
	"sync"
	"testing"
)

func TestSync(t *testing.T) {

	assertCorrectValue := func(t *testing.T, got *Counter, want int) {
		t.Helper()
		if got.Value() != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 3
		assertCorrectValue(t, counter, want)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCorrectValue(t, counter, wantedCount)

	})

}
