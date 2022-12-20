package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jonas")
	want := "Hello, Jonas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
