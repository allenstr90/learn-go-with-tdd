package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://isos.uci.cu",
		"http://google.com",
	}

	want := map[string]bool{
		"http://isos.uci.cu": true,
		"http://google.com":  false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v but got %v", want, got)
	}
}

func mockWebsiteChecker(str string) bool {
	return str == "http://isos.uci.cu"
}

func slowDownChecker(str string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}


func BenchmarkWebsiteChecker(b *testing.B) {
	urls:= make([]string, 100)
	for i:=0; i< len(urls); i++{
		urls[i] = "a url "
	}
	b.ResetTimer()
	for i:=0; i< b.N; i++{
		CheckWebsites(slowDownChecker, urls)
	}
}