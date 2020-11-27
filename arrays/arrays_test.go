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

	t.Run("sum each array", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertCorrectMessage(t, got, want)
	})

}
func TestSumAllTails(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of tail[1:] from array", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 4, 5}, []int{0, 9, 0, 1})
		want := []int{11, 10}
		assertCorrectMessage(t, got, want)
	})

	t.Run("test empty array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 9, 5})
		want := []int{0, 14}
		assertCorrectMessage(t, got, want)
	})

}
