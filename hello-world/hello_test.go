package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("paco")
	want := "hello paco !!!"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}