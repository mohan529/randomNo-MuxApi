package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", Output)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Output(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Random Number is: %v", randomeNumbers())
}

func randomeNumbers() int {
	return rand.Intn(1000)
}
