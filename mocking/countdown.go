/**
We have covered:
- Without mocking important areas of your code will be untested.
- Test the behaviour of your code rather than the implementation.
- Spies: `test-double`, replace production objects for testing purpose.
*/

package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"time"
)

const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type SpyCoutdownOperation struct {
	Calls []string
}

const (
	sleep = "sleep"
	write = "write"
)

func (s *SpyCoutdownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCoutdownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// A simple iterator that yield numbers in reverse order
func countdownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := range countdownFrom(3) {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}