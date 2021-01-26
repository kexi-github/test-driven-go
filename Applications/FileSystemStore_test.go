package Applications_test

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"tests/Applications"
)

func TestFileSystemStore(t *testing.T){

	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name":"Cleo","Wins":10},
		{"Name":"Chris","Wins":33}]`)
		defer cleanDatabase()

		store := Applications.NewFileSystemStore(database)

		got := store.GetLeague()

		want := []Applications.Player{
			{"Cleo",10},
			{"Chris",33},
		}

		if !reflect.DeepEqual(got,want){
			t.Errorf("got %v,want %v",got,want)
		}

		// read again
		got = store.GetLeague()
		if !reflect.DeepEqual(got,want){
			t.Errorf("got %v,want %v",got,want)
		}
	})

	t.Run("get player score", func(t *testing.T) {
		database,cleanDatabase := createTempFile(t,`[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store := Applications.NewFileSystemStore(database)
		got := store.GetPlayerScore("Chris")
		want := 33
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store := Applications.NewFileSystemStore(database)
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store := Applications.NewFileSystemStore(database)
		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store := Applications.NewFileSystemStore(database)
		got := store.GetLeague()
		want := []Applications.Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		if !reflect.DeepEqual(got,want){
			t.Errorf("got %v,want %v",got,want)
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
