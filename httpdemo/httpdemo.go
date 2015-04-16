package main

import (
	"fmt"
	"log"
	"net/http"
)

// Must implement http.Handler
type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello from server, desu!")
	fmt.Fprintln(w, "Your method was ", r.Method)
}

func main() {
	var h Hello // creating new server

	err := http.ListenAndServe("localhost: 4000", h)

	if err != nil {
		log.Fatal(err)
	}
}
