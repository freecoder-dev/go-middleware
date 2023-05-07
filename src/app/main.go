package main

import (
	"fmt"
	"log"
	"net/http"
)

var Revision string

type fctHandler http.HandlerFunc

func reportLog(f fctHandler) fctHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello freecoders!")
}

func HiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi freecoders!")
}

func main() {
	http.HandleFunc("/hi", reportLog(HiHandler))
	http.HandleFunc("/hello", reportLog(HelloHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
