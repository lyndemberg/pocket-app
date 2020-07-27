package controller

import (
	"encoding/json"
	"fmt"
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
	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)
	user, err := userRepository.FindByID(id)

	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func userDeleteAction(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)
	err := userRepository.DeleteByID(id)

	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, "User successfully deleted")
	}

	w.Header().Add("Content-Type", "text/plain")
}

func userCreateAction(w http.ResponseWriter, r *http.Request) {
	var userRequest model.User
	json.NewDecoder(r.Body).Decode(&userRequest)

	user, err := userRepository.Create(userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("error", err.Error())
	} else {
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusCreated)
	}
}

func userUpdateAction(w http.ResponseWriter, r *http.Request) {
	var userRequest model.User
	json.NewDecoder(r.Body).Decode(&userRequest)

	user, err := userRepository.Update(userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("error", err.Error())
	} else {
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusCreated)
	}
}
