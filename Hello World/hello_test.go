package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Wolle")
	want := "Hello Wolle"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
