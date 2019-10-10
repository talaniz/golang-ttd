package mocks

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper represents a sleep interval
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is the sleepr to use in the main implementation
type DefaultSleeper struct{}

// Sleep implements a wait
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown prints from 3 to 1
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}
