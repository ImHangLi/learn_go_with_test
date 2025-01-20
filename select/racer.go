/**
We have covered:
- select
	- Helps you wait on multiple channels.
	- Sometimes you'll want to include `time.After` in one of your cases to prevent your system blocking forever.
- httptest
	- A convenient way of creating test servers so you can have reliable and controllable tests.
	- Uses the same interfaces as the "real" net/http servers which is consistent and less for you to learn.
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
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
	ch := make(chan struct{})

	go func ()  {
		http.Get(url)
		close(ch)
	} ()

	return ch
}