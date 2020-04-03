package interation

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {

	ErrorMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if strings.Compare(got, want) != 0 {
			t.Errorf("[FAILED]\ngot => %q\nwant => %q", got, want)
		}
	}

	t.Run("Say 'a' 5 times to user", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		ErrorMessage(t, got, want)
	})

	t.Run("Say 'a' 3 times ", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"

		ErrorMessage(t, got, want)
	})

	t.Run("Say 'a' 0 times ", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""

		ErrorMessage(t, got, want)
	})

	t.Run("Say 'a' -1 times ", func(t *testing.T) {
		got := Repeat("a", -1)
		want := ""

		ErrorMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
