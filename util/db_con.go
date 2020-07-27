package util

import (
	"database/sql"
	"fmt"
	"os"

	//import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var user string
var password string
var database string

func init() {
	godotenv.Load(".env")
	user = os.Getenv("db_user")
	password = os.Getenv("db_password")
	database = os.Getenv("db_name")
}

// GetConnection TODO
func GetConnection() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@/%s", user, password, database)
	return sql.Open("mysql", dataSourceName)
}
