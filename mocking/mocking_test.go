package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

type Spysleeper struct {
	Calls int
}

func (s Spysleeper) Sleep() {
	s.Calls++
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

func TestMock(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Mock Countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		assertCorrectMessage(t, got, want)
	})

	t.Run("Real Countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		defaultSleeper := &DefaultSleeper{}

		Countdown(buffer, defaultSleeper)

		got := buffer.String()
		want := `3
2
1
Go!`
		assertCorrectMessage(t, got, want)
	})

	t.Run("Sleep4Printing Countdown", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

}
