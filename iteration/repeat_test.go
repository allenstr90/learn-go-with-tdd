package main

import "testing"

func TestRepeat(t *testing.T) {
	rep := Repeat("a")
	exp := "aaaaa"

	if rep != exp {
		t.Errorf("exp %q but got %q", exp, rep)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a");
	}
}
