package Integers_test

import (
	"fmt"
	"testing"
	"tests/Integers"
)

func TestAdder(t *testing.T){

	sum := Integers.Add(2,2)
	expected := 4

	if sum != expected{
		t.Errorf("expected %d but got %d",expected,sum)
	}
}

func ExampleAdd() {

	sum := Integers.Add(1,5)
	fmt.Println(sum)
	// Output: 6
}
