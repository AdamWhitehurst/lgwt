package integers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		want := 4
		got := Add(2, 2)

		if want != got {
			t.Errorf("expected '%d' but got '%d'", want, got)
		}
	})
}
