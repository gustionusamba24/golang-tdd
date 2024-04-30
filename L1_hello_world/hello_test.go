package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("gusti")
	want := "Hello, gusti"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
