package racer

import (
	"fmt"
	"net/http"
	"time"
)

const defaultTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a,b, defaultTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timed out (%v) waiting for %s and %s", timeout, a, b)		
	}
}

func checkResponseTime(url string) time.Duration{
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}

func ping(url string) chan struct{}{
	c := make(chan struct{})
	go func()  {
		http.Get(url)
		close(c)
	}()
	return c
}