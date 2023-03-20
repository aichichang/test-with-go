package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("blah")
		want := ErrNotFound.Error()

		// We are also protecting assertStrings with if statement
		// to ensure we don't call .Error() on nil.
		if err == nil {
			t.Fatal("expect an error")
		}

		// Errors can be converted to a string with the .Error()
		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {

	t.Run("add", func(t *testing.T) {
		word := "test"
		def := "this is a test"
		dictionary := Dictionary{}

		dictionary.Add(word, def)

		assertDef(t, dictionary, word, def)

	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is a test"
		dictionary := Dictionary{word: def}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDef(t, dictionary, word, def)
	})
}

func assertDef(t testing.TB, dictionary Dictionary, word, def string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Errorf("should find added word %v", err)
	}

	assertStrings(t, got, def)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
