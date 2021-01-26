package main

import (
	"log"
	"net/http"
	"os"
	"tests/Applications"
)

const dbFileName = "game.db.json"

func main() {

	db,err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	store := Applications.NewFileSystemStore(db)

	server := Applications.NewPlayerServer(store)
	if err:=http.ListenAndServe(":5000",server);err != nil{
		log.Fatal("ListenAndServe error")
	}
}
