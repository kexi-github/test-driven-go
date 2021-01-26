package main_test

import (
	"bytes"
	"testing"
	"tests/DI"
)

func TestGreet(t *testing.T){

	buffer := bytes.Buffer{}
	main.Greet(&buffer,"Chris")

	got := buffer.String()
	want := "Hello,Chris"

	if got != want{
		t.Errorf("got %s want %s",got,want)
	}
}
