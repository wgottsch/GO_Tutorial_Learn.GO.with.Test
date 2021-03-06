package maps

import (
	"github.com/pkg/errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"                      //key
		definition := "this is just a test" //value

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"                      //key
		definition := "this is just a test" //value
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("Dictonary: Change value on existing Key", func(t *testing.T) {

		word := "test"                      //key
		definition := "this is just a test" //value
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Dictonary: Change value on non existing Key", func(t *testing.T) {
		word := "test"                      //key
		definition := "this is just a test" //value
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrNotFound)
	})

}

func TestDelete(t *testing.T) {
	t.Run("Dictonary: Delete value on existing Key", func(t *testing.T) {

		word := "test"                      //key
		definition := "this is just a test" //value
		dictionary := Dictionary{word: definition}

		delWord := "test"
		err := dictionary.Delete(delWord)

		assertError(t, err, nil)
	})

	t.Run("Dictonary: Delete value on non existing Key", func(t *testing.T) {
		word := "test"                      //key
		definition := "this is just a test" //value
		dictionary := Dictionary{word: definition}

		delWord := "ups"
		err := dictionary.Delete(delWord)

		assertError(t, err, ErrNotFound)
	})

}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if errors.Cause(got) != want {
		t.Errorf("got error %q want %q", got, want)
	}
	if errors.Cause(got) == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error")
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
