package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.search("test")
		expected := "this is just a test"
		assertString(t, got, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.search("unknown")
		want := notFoundError

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertString(t, err.Error(), want.Error())
	})

}
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictionary.add(word, definition)
		got, err := dictionary.search(word)
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		assertString(t, got, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.add(word, "new test")

		assertString(t, err.Error(), wordExistsError.Error())
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update #table", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "Updated"
		dictionary.update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})
}

func ()  {
	
}

func assertString(t *testing.T, got string, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("Got: %q Expected: %q", got, expected)
	}
}
func assertDefinition(t *testing.T, dictionary Dictionary, key string, expected string) {
	t.Helper()
	got, err := dictionary.search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, got, expected)
}
