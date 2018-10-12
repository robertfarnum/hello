package main

import "testing"

func TestHello(t *testing.T) {
	// Assert on correct message function declaration
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	// Test saying hello to people
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", English)
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	// Test empty string default to 'World'
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", English)
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	// Test Spanish'
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", Spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	// Test French'
	t.Run("in Frensh", func(t *testing.T) {
		got := Hello("Robert", French)
		want := "Bonjour, Robert"
		assertCorrectMessage(t, got, want)
	})
}
