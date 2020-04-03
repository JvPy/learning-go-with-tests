package integers

import "testing"

func TestAdder(t *testing.T) {
	ErrorMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("[FAILED]\ngot => %d\nwant => %d", got, want)
		}
	}

	t.Run("Should sum 2 integers", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		ErrorMessage(t, got, want)
	})

	t.Run("Should sum 2 integers", func(t *testing.T) {
		got := Add(2, -2)
		want := 0

		ErrorMessage(t, got, want)
	})
}
