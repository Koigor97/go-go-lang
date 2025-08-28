package main

import "testing"

func TestGreetingWorld(t *testing.T) {
	want := "Hello world"

	got := greetingsWorld()

	if got != want {
		// mark as test failed
		t.Errorf("expected: %q. got: %q", want, got)
	}
}
