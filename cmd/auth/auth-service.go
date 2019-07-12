package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello from login"))
	fmt.Print(r)
}
