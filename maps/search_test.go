package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "map test"}

	got := Search(dict, "test")
	want := "map test"

	assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestDict(t *testing.T) {
	dict := Dict{}
	dict.Add("test", "map test")

	t.Run("add new word", func(t *testing.T) {
		err := dict.Add("new", "word")
		assertError(t, err, nil)
	})

	t.Run("duplicate word", func(t *testing.T) {
		err := dict.Add("new", "word")
		assertError(t, err, WordAlreadyExistsException)
	})

	t.Run("found word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "map test"

		assertString(t, got, want)
	})

	t.Run("no found", func(t *testing.T) {
		_, err := dict.Search("test2")

		assertError(t, err, WordNotFoundException)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q but want %q", got, want)
	}
}
