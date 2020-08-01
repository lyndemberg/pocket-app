package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyndemberg/pocket-app/controller"
	"github.com/lyndemberg/pocket-app/security"
)

func main() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	controller.NewLoginController().Handle(router.PathPrefix("/login").Subrouter())
	controller.NewUserController().Handle(apiRouter.PathPrefix("/users").Subrouter())

	router.Use(security.AuthorizationMiddleware)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
