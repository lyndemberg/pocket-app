package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken TODO comment
func GenerateToken(userID string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
