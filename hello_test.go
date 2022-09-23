package main

import "testing"

func TestHelloX(t *testing.T) {
	got := Hello("x")
	want := "Hello, x"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
