package main

import "testing"

func TestAdder(t *testing.T){
	got := Adder(1,1)
	want := 2

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}