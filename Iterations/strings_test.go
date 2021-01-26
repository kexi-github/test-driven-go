package Iterations

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleCompare()  {

	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
	// Output: -1
	//0
	//1
}

func ExampleContains()  {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	// Output: true
	//false
	//true
	//true
}

func TestContains(t *testing.T)  {

	got := strings.Contains("seafood", "foo")
	want := true

	if got != want{
		t.Errorf("got %v want %v",got,want)
	}
}
