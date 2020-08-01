package controller

import (
	"net/http"

	"github.com/lyndemberg/pocket-app/security"
	"github.com/lyndemberg/pocket-app/util"

	"github.com/gorilla/mux"
	"github.com/lyndemberg/pocket-app/repository"
)

//LoginController TODO
type LoginController struct {
	userRepository *repository.UserRepository
}

// NewLoginController TODO comment
func NewLoginController() *LoginController {
	l := new(LoginController)
	l.userRepository = repository.NewUserRepository()
	return l
}

//Handle TODO
func (control LoginController) Handle(route *mux.Router) {
	route.HandleFunc("", control.executeLogin).Methods("POST")
}

func (control LoginController) executeLogin(w http.ResponseWriter, r *http.Request) {
	usernameRequest := r.FormValue("username")
	passwordRequest := r.FormValue("password")

	user, err := control.userRepository.FindByUsername(usernameRequest)
	_, isLogicError := err.(util.LogicError)

	if err != nil && !isLogicError {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("error", "There was a problem signing in")
	} else if err != nil && isLogicError {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("error", "Check your credentials")
	} else {
		// user exists
		passwordIsCorrect := security.CheckPassword(passwordRequest, user.Password)
		if !passwordIsCorrect {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("error", "Check your credentials")
		} else {
			// generate token
			token, _ := security.GenerateToken(user.ID)
			w.Write([]byte(token))
			w.WriteHeader(http.StatusOK)
		}
	}
}
