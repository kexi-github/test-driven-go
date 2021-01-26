package Select_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"tests/Select"
	"time"
)

func TestRacer(t *testing.T)  {

	slowServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(20 * time.Millisecond)
		writer.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Select.Racer(slowURL,fastURL)

	if got != want{
		t.Errorf("got %s,want %s",got,want)
	}
}
