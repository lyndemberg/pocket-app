package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/lyndemberg/pocket-app/controllers"
)

func main() {
	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()

	controller.UserController(router.PathPrefix("/users").Subrouter())

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
