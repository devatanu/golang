package hello

import "testing"

func TestGetGreeting(t *testing.T) {
	got := getGreeting()
	want := "Hello world"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
