package Iterations_test

import (
	"testing"
	"tests/Iterations"
)

func TestRepeat(t *testing.T){

	repeated := Iterations.Repeat("a")
	expected := "aaaaaa"

	if repeated != expected{
		t.Errorf("expected %s but got %s",expected,repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {

	for i :=0;i<b.N;i++{
		Iterations.Repeat("a")
	}
}
