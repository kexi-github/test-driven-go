package Maps_test

import (
	"testing"
	"tests/Maps"
)

func TestSearch(t *testing.T){

	dictionary := Maps.Dictionary{}
	dictionary.Add("test","this is just a test")

	t.Run("known word", func(t *testing.T) {
		got,_ := dictionary.Search("test")
		want := "this is just a test"

		if got != want{
			t.Errorf("got %s want %s given,%s",got,want,"test")
		}
	})
	t.Run("unknown word", func(t *testing.T) {
		got,err := dictionary.Search("unknow")
		want := "could not find the word you were looking for"

		if err == nil{
			t.Errorf("got %s want %s expected an error",got,want)
		}
	})
	t.Run("exists word", func(t *testing.T) {
		err := dictionary.Add("test","value")

		if err == nil{
			t.Fatalf("expected an error")
		}
	})
	t.Run("update word", func(t *testing.T) {

		_ = dictionary.Update("test","value")

		got,_ := dictionary.Search("test")
		want := "value"

		if got != want{
			t.Errorf("got %s want %s given,%s",got,want,"test")
		}
	})
	t.Run("delete word", func(t *testing.T) {

		_ = dictionary.Delete("test")

		got,_ := dictionary.Search("test")
		want := ""

		if got != want{
			t.Errorf("got %s want %s given,%s",got,want,"test")
		}
	})
}
