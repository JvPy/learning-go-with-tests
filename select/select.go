package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
func ping(url string) chan struct{} {
	/*
		Using a empty struct because its the smallest data type
		As we only need to close the channel to signal the return
		we dont need to alocate memory to a bool
	*/

	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
