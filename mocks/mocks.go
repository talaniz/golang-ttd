package mocks

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// ConfigurableSleeper is a sleeper with configs
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep will pause execution for the defined duration
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Sleeper represents a sleep interval
type Sleeper interface {
	Sleep()
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
