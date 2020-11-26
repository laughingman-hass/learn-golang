package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Hassan")
    want := "Hello, Hassan"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
