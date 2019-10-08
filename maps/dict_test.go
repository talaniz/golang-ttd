package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	assertStrings := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected an error")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		if err != nil {
			t.Fatal("did not expect an error, but one was generated")
		}

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		if err != nil {
			t.Fatal("did not expect an error, but one was generated")
		}

		assertDefinition(t, dictionary, word, definition)

		err = dictionary.Add(word, definition)

		assertError(t, err, ErrAlreadyExists)

	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	if definition != got {
		t.Errorf("got %q, wanted %q", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
