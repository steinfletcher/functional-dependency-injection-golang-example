package main

import (
	"github.com/labstack/gommon/log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/steinfletcher/injection-play/user"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/user/{id}", user.Module())
	log.Fatal(http.ListenAndServe(":8080", router))
}
