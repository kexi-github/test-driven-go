package Goroutines_test

import (
	"reflect"
	"testing"
	"tests/Goroutines"
	"time"
)

func mockWebsiteChecker(url string) bool{

	if url == "waat://furhurterwe.geds"{
		return false
	}
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T){

	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	actualResults := Goroutines.CheckWebsites(mockWebsiteChecker,websites)

	want := len(websites)
	got := len(actualResults)

	if want != got{
		t.Fatalf("want %v,got %v",want,got)
	}

	expectedResults := map[string]bool{
		"http://google.com":true,
		"http://blog.gypsydave5.com":true,
		"waat://furhurterwe.geds":false,
	}

	if !reflect.DeepEqual(expectedResults,actualResults){
		t.Fatalf("want %v,got %v",expectedResults,actualResults)
	}
}

func BenchmarkCheckWebsites(b *testing.B){

	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	for i:=0;i<b.N;i++{
		Goroutines.CheckWebsites(mockWebsiteChecker,websites)
	}
}
