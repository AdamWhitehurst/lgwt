package lgwt

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Adam", "")
		want := "Hello, Adam"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when greetee is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Dago", "Spanish")
		want := "Hola, Dago"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Dago", "French")
		want := "Bonjour, Dago"

		assertCorrectMessage(t, got, want)
	})
}
