package mock

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("print from 3 to Go!!!", func(t *testing.T) {
		buff := &bytes.Buffer{}
		Countdown(buff, &SpyCountdownOps{})

		got := buff.String()
		want := "3\n2\n1\nGo!!!"

		if want != got {
			t.Errorf("want %s but got %s", want, got)
		}
	})

	t.Run("sleep before print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOps{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write, sleep, write, sleep, write, sleep, write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v but got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

// mock classes
type SpyCountdownOps struct {
	Calls []string
}

const sleep = "sleep"
const write = "write"

func (s *SpyCountdownOps) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOps) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
