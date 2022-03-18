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
		got := Hello("paco", "en")
		want := "hello paco !!!"

		assertCorrectMesg(t, got, want)
	})

	t.Run("saying hello world when empty name", func(t *testing.T) {
		got := Hello("", "en")
		want := "hello world !!!"

		assertCorrectMesg(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got:= Hello("paco", "es")
		want := "hola paco !!!"
		assertCorrectMesg(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got:= Hello("paco", "fr")
		want := "bonjour paco !!!"
		assertCorrectMesg(t, got, want)
	})
}
