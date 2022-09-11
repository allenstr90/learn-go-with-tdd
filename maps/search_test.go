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
		assertDefinition(t, dict, "new", "word")
	})

	t.Run("duplicate word", func(t *testing.T) {
		err := dict.Add("new", "word")
		assertError(t, err, WordAlreadyExistsException)
		assertDefinition(t, dict, "new", "word")
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

func TestUpdate(t *testing.T) {
	key := "test"
	value := "map old"
	dict := Dict{}
	dict.Add(key, value)

	newValue := "map new"
	dict.Update(key, newValue)
	assertDefinition(t, dict, key, newValue)
}

func TestDelete(t *testing.T) {
	key := "test"
	value := "value"
	dict := Dict{}
	dict.Add(key, value)

	dict.Delete(key)

	_, err := dict.Search(key)
	assertError(t, err, WordNotFoundException)
}

func assertDefinition(t testing.TB, dict Dict, key, want string) {
	t.Helper()

	got, err := dict.Search(key)

	if err != nil {
		t.Fatal("should find the word:", err)
	}

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q but want %q", got, want)
	}
}
