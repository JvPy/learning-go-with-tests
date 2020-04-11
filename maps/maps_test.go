package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	assertStrings := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\ngot: %s,\nwant: %s", got, want)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got: %q,\nwant: %q", got, want)
		}
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unkown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	assertDefinition := func(t *testing.T, dictionary Dictionary, word, definition string) {
		t.Helper()

		got, err := dictionary.Search(word)
		if err != nil {
			t.Fatal("should find added word:", err)
		}

		if definition != got {
			t.Errorf("got %q want %q", got, definition)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got: %q,\nwant: %q")
		}

		if got == nil {
			if want == nil {
				return
			}
			t.Fatal("expected to get an error.")
		}
	}

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "definition"

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrNotFound)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}
