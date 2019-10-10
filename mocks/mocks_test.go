package mocks

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("test coundown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
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

// CountdownOperationsSpy tracks operations for Countdown
type CountdownOperationsSpy struct {
	Calls []string
}

// Sleep implements sleep for CountdownOperationsSpy
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write implements the writer interface for the CountdownOperationsSpy
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// SpyTime is a configuration duration
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep implements the time.Sleep interface for SpyTime
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
