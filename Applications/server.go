package Applications

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PlayStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	Store PlayStore
	http.Handler
}

func NewPlayerServer(store PlayStore) *PlayerServer {
	p := new(PlayerServer)

	p.Store = store

	router := http.NewServeMux()

	router.Handle("/league",http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("content-type", "application/json")

		leagueTable := p.Store.GetLeague()

		json.NewEncoder(writer).Encode(leagueTable)
	}))

	router.Handle("/players/",http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		player := request.URL.Path[len("/players/"):]

		if request.Method == http.MethodPost{
			writer.WriteHeader(http.StatusAccepted)
			p.Store.RecordWin(player)
		}

		if request.Method == http.MethodGet{
			if p.Store.GetPlayerScore(player) == 0{
				writer.WriteHeader(http.StatusNotFound)
			}else{
				writer.WriteHeader(http.StatusOK)
			}
			fmt.Fprint(writer,p.Store.GetPlayerScore(player))
		}
	}))

	p.Handler = router

	return p
}
