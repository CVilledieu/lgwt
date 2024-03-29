package main

import "testing"

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
	word := "test"
	definition := "this is just a test"
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("word exists", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {
	word := "test"
	newDefinition := "new updated definiton"
	oldDefinition := "this is just a test"
	dictionary := Dictionary{word: oldDefinition}

	t.Run("update word", func(t *testing.T) {
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)

	})

	t.Run("new word", func(t *testing.T) {
		newDictionary := Dictionary{}
		err := newDictionary.Update(word, newDefinition)

		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is a test"
	dictionary := Dictionary{word: definition}

	t.Run("word deleted", func(t *testing.T) {
		dictionary.Delete(word)
		assertDelete(t, dictionary, word)
	})

	t.Run("word not found", func(t *testing.T) {
		err := dictionary.Delete(word)

		assertError(t, err, ErrDeleteWord)
	})

}

func assertDelete(t testing.TB, dictionary Dictionary, word string) {
	t.Helper()

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("found %q, but should have been deleted", err)
	}

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q wanted %q was given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
