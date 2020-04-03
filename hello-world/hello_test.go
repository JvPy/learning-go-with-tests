package main

import "testing"

func TestHello(t *testing.T) {

	ErrorMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("[FAILED]\ngot => %q\nwant => %q", got, want)
		}
	}

	t.Run("Say hello to user", func(t *testing.T) {
		got := SayHello("Joao", "")
		want := "Hello, Joao"

		ErrorMessage(t, got, want)
	})

	t.Run("Say 'Hello, world' when empty string", func(t *testing.T) {
		got := SayHello("", "")
		want := "Hello, world"

		ErrorMessage(t, got, want)
	})

	t.Run("Say hello to spanish user", func(t *testing.T) {
		got := SayHello("Arturo", "Spanish")
		want := "Hola, Arturo"
		ErrorMessage(t, got, want)
	})

	t.Run("Say hello to portuguese user", func(t *testing.T) {
		got := SayHello("Joao", "Portuguese")
		want := "Olá, Joao"
		ErrorMessage(t, got, want)
	})
}
