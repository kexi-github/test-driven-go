package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"tests/Applications"
)

func TestCLI(t *testing.T){

	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name":"Cleo","Wins":10},
		{"Name":"Chris","Wins":33}]`)
		defer cleanDatabase()

		store := Applications.NewFileSystemStore(database)

		in := strings.NewReader("Cleo wins\n")

		cli := NewCLI(store,in)
		cli.PlayPoker()
		got := cli.playStore.GetPlayerScore("Cleo")
		want := 11

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removeFile
}
