package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer,str string) {

	fmt.Fprintf(writer,"Hello,%s",str)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
