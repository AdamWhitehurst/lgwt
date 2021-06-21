package lgwt

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Adam")
	want := "Hello, Adam"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
