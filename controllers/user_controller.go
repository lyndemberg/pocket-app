package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/lyndemberg/pocket-app/models"

	"github.com/gorilla/mux"
	"github.com/lyndemberg/pocket-app/repository"
)

var userRepository = repository.NewUserRepository()

// UserController TODO
func UserController(subRouter *mux.Router) {
	subRouter.HandleFunc("", userListAction).Methods("GET")
	subRouter.HandleFunc("", userCreateAction).Methods("POST")
	subRouter.HandleFunc("/{id}", userDetailsAction).Methods("GET")
	subRouter.HandleFunc("/{id}", userUpdateAction).Methods("PUT")
	subRouter.HandleFunc("/{id}", userDeleteAction).Methods("DELETE")
}

func userListAction(w http.ResponseWriter, r *http.Request) {
	userList := userRepository.FindAll()
	json.NewEncoder(w).Encode(userList)
}

func userDetailsAction(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idString, _ := strconv.Atoi(id)
	user := userRepository.FindByID(idString)
	json.NewEncoder(w).Encode(user)
}

func userDeleteAction(w http.ResponseWriter, r *http.Request) {

}

func userCreateAction(w http.ResponseWriter, r *http.Request) {
	var userRequest model.User
	json.NewDecoder(r.Body).Decode(&userRequest)

	user, err := userRepository.Create(userRequest)
	if err == nil {
		json.NewEncoder(w).Encode(user)
	}
}

func userUpdateAction(w http.ResponseWriter, r *http.Request) {

}
