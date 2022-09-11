package di

import (
	"bytes"
	"testing"
)

func TestDi(t *testing.T) {
	buff := bytes.Buffer{}
	Greet(&buff, "Allen")

	got := buff.String()
	want := "Hello Allen"

	if want != got {
		t.Errorf("got %s but want %s", got, want)
	}
}
