package Mocking_test

import (
	"bytes"
	"testing"
	"tests/Mocking"
)

func TestCountdown(t *testing.T){

	buffer := &bytes.Buffer{}

	Mocking.Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want{
		t.Errorf("got %s want %s",got,want)
	}
}
