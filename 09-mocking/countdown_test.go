package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	actual := buffer.String()
	expected := "3\n2\n1\nGo!"

	if actual != expected {
		t.Errorf("actual %q, expected %q", actual, expected)
	}

    if spySleeper.Calls != 4 {
        t.Errorf("not enough calls to sleeper, expected 4, actual %d", spySleeper.Calls)
    }
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
