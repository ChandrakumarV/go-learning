package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Mod")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	http.ListenAndServe(":5100", r)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello 0</h1>"))

}
