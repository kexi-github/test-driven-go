package Applications_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"tests/Applications"
)

func TestPlayerServer(t *testing.T){

	store := Applications.StubPlayserStore{
		map[string]int{
			"Pepper":20,
			"Floyd":10,
		},
	}
	playerserver := Applications.NewPlayerServer(&store)

	t.Run("return Pepper's score", func(t *testing.T) {
		request,_ := http.NewRequest(http.MethodGet,"/players/Pepper",nil)
		response := httptest.NewRecorder()

		playerserver.ServeHTTP(response,request)

		got := response.Body.String()
		want := "20"

		if got != want{
			t.Errorf("got %s,want %s",got,want)
		}
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()
		playerserver.ServeHTTP(response,request)
		got := response.Body.String()
		want := "10"
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request,_ := http.NewRequest(http.MethodGet, "/players/Apollo", nil)
		response := httptest.NewRecorder()
		playerserver.ServeHTTP(response,request)
		got := response.Code
		want := http.StatusNotFound
		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Peppers", nil)
		response := httptest.NewRecorder()
		playerserver.ServeHTTP(response,request)
		got := response.Code
		want := http.StatusAccepted
		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}

		request1,_ := http.NewRequest(http.MethodGet, "/players/Peppers", nil)
		response1 := httptest.NewRecorder()
		playerserver.ServeHTTP(response1,request1)
		got1 := response1.Body.String()
		want1 := "1"
		if got1 != want1 {
			t.Errorf("got '%s', want '%s'", got1, want1)
		}
	})
}

func TestLeague(t *testing.T){

	t.Run("it returns 200 on /league", func(t *testing.T) {

		store := Applications.StubPlayserStore{
			map[string]int{
				"Pepper":20,
				"Floyd":10,
			},
		}
		playerserver := Applications.NewPlayerServer(&store)

		request,_ := http.NewRequest(http.MethodGet,"/league",nil)
		response := httptest.NewRecorder()

		playerserver.ServeHTTP(response,request)

		var got []Applications.Player
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil{
			t.Fatalf("Unable to parse response from server")
		}
		want := []Applications.Player{
			{"Pepper",20},
			{"Floyd",10},
		}

		if !reflect.DeepEqual(got,want){
			t.Errorf("got %v,want %v",got,want)
		}
	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {

		wantedLeague := []Applications.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := Applications.StubPlayserStore{
			map[string]int{
				"Cleo": 32,
				"Chris": 20,
				"Tiest": 14,
			},
		}
		playerserver := Applications.NewPlayerServer(&store)

		request,_ := http.NewRequest(http.MethodGet,"/league",nil)
		response := httptest.NewRecorder()

		playerserver.ServeHTTP(response,request)

		var got []Applications.Player
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil{
			t.Fatalf("Unable to parse response from server")
		}

		if response.Header().Get("content-type") != "application/json" {
			t.Errorf("response did not have content-type of application/json, got %v", response.Header())
		}

		if !reflect.DeepEqual(got,wantedLeague){
			t.Errorf("got %v,want %v",got,wantedLeague)
		}
	})
}
