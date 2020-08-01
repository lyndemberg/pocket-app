package security

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// AuthorizationMiddleware TODO comment
func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
		} else {
			_, err := verifyAuthorization(w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
			}
			next.ServeHTTP(w, r)
		}
	})
}

func verifyAuthorization(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	token := extractToken(r)
	if token != "" {
		tokenInstance, err := jwt.Parse(token, keyFunction)
		if err != nil {
			return nil, errors.New("Invalid token")
		}
		return tokenInstance, nil
	}
	return nil, errors.New("Token not present")
}

func keyFunction(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(SecretKey), nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
