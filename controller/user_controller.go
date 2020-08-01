package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/lyndemberg/pocket-app/model"
	"github.com/lyndemberg/pocket-app/security"

	"github.com/gorilla/mux"
	"github.com/lyndemberg/pocket-app/repository"
)

//UserController TODO type comment
type UserController struct {
	userRepository *repository.UserRepository
}

// NewUserController TODO comment
func NewUserController() *UserController {
	u := new(UserController)
	u.userRepository = repository.NewUserRepository()
	return u
}

// Handle Method
func (control UserController) Handle(subRouter *mux.Router) {
	subRouter.HandleFunc("", control.userListAction).Methods("GET")
	subRouter.HandleFunc("", control.userCreateAction).Methods("POST")
	subRouter.HandleFunc("/{id}", control.userDetailsAction).Methods("GET")
	subRouter.HandleFunc("/{id}", control.userUpdateAction).Methods("PUT")
	subRouter.HandleFunc("/{id}", control.userDeleteAction).Methods("DELETE")
}

func (control UserController) userListAction(w http.ResponseWriter, r *http.Request) {
	userList := control.userRepository.FindAll()
	json.NewEncoder(w).Encode(userList)
}

func (control UserController) userDetailsAction(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)
	user, err := control.userRepository.FindByID(id)

	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func (control UserController) userDeleteAction(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)
	err := control.userRepository.DeleteByID(id)

	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, "User successfully deleted")
	}

	w.Header().Add("Content-Type", "text/plain")
}

func (control UserController) userCreateAction(w http.ResponseWriter, r *http.Request) {
	var userRequest model.User
	json.NewDecoder(r.Body).Decode(&userRequest)

	hashedPassword, errHashPassword := security.PasswordToHash(userRequest.Password)
	if errHashPassword != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("error", "There was a problem processing the user registration")
	} else {
		userRequest.Password = hashedPassword
		user, err := control.userRepository.Create(userRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("error", err.Error())
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(user)
		}
	}
}

func (control UserController) userUpdateAction(w http.ResponseWriter, r *http.Request) {
	var userRequest model.User
	json.NewDecoder(r.Body).Decode(&userRequest)

	user, err := control.userRepository.Update(userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("error", err.Error())
	} else {
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusCreated)
	}
}
