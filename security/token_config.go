package security

import (
	"os"

	"github.com/joho/godotenv"
)

// SecretKey is the key used to generate the tokens
// of the users of the application during the login process
var SecretKey string

func init() {
	godotenv.Load(".env")
	SecretKey = os.Getenv("jwt_key")
}
