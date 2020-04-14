package main

import (
	"bytes"
	"testing"
)

func TestGree(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Joao")

	got := buffer.String()
	want := "Hello, Joao"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
