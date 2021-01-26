package main_test

import (
	"testing"
	"tests/Hello"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T,got,want string) {
		t.Helper()
		if got != want{
			t.Errorf("got %s want %s",got,want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := main.Hello("Chris","")
		want := "Hello,Chris"

		assertCorrectMessage(t,got,want)
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := main.Hello("","")
		want := "Hello,world"

		assertCorrectMessage(t,got,want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := main.Hello("Elodie","Spanish")
		want := "Hola,Elodie"

		assertCorrectMessage(t,got,want)
	})

	t.Run("in French", func(t *testing.T) {
		got := main.Hello("Full","French")
		want := "Bonjour,Full"

		assertCorrectMessage(t,got,want)
	})
}
