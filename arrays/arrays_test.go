package arrays

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertCorrectMessage(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	got := SumAll([]int{1, 2}, []int{0, 9})
	//want := "bobby"  (eigentlich compiliert go das im tutorial, durchs assert fällt es hier aber auf :-)
	want := []int{3, 9}

	assertCorrectMessage(t, got, want)
}
