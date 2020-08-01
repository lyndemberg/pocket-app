package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/lyndemberg/pocket-app/controllers"
	"github.com/lyndemberg/pocket-app/security"
)

func main() {
	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()

	router.Use(security.AuthorizationMiddleware)

	controller.NewUserController().Handle(router.PathPrefix("/users").Subrouter())
	controller.NewLoginController().Handle(router.PathPrefix("/login").Subrouter())

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
