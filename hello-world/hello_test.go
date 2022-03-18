package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMesg := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("paco")
		want := "hello paco !!!"

		assertCorrectMesg(t, got, want)
	})

	t.Run("saying hello world when empty name", func(t *testing.T) {
		got := Hello("")
		want := "hello world !!!"

		assertCorrectMesg(t, got, want)
	})
}
