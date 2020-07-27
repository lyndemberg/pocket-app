package repository

import (
	"database/sql"
	"errors"
	"fmt"

	model "github.com/lyndemberg/pocket-app/models"
	util "github.com/lyndemberg/pocket-app/util"
)

//UserRepository TODO
type UserRepository struct {
	connection *sql.DB
}

//NewUserRepository TODO
func NewUserRepository() *UserRepository {
	u := new(UserRepository)
	db, err := util.GetConnection()
	u.connection = db
	fmt.Print(db)
	fmt.Print(err)
	if err != nil {
		u.connection = nil
	}
	return u
}

//FindAll TODO
func (urepo UserRepository) FindAll() []model.User {
	var list = make([]model.User, 0)

	if urepo.connection == nil {
		return list
	}

	rows, err := urepo.connection.Query("SELECT id, name, email, username, password FROM users")

	if err != nil {
		return list
	}

	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Username, &u.Password)
		list = append(list, u)
	}

	defer rows.Close()
	defer urepo.connection.Close()
	return list
}

//FindByID TODO comment
func (urepo UserRepository) FindByID(id int) model.User {
	var u model.User

	if urepo.connection != nil {
		query := "SELECT id, name, email, username, password FROM users WHERE id = ?"
		urepo.connection.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email, &u.Username, &u.Password)
	}

	return u
}

//Create TODO comment
func (urepo UserRepository) Create(user model.User) (model.User, error) {
	if urepo.connection != nil {
		sqlInsert := "INSERT INTO users (name, email, username, password) VALUES (?, ?, ?, ?)"
		result, err := urepo.connection.Exec(sqlInsert, user.Name, user.Email, user.Username, user.Password)

		lastID, _ := result.LastInsertId()
		return urepo.FindByID(int(lastID)), err
	}

	return user, errors.New("No connection with database")
}

//Update TODO comment
// func (urepo UserRepository) Update(user model.User) (model.User, error) {
// 	if urepo.connection != nil {
// 		sqlUpdate := "UPDATE users SET name = ?, email = ?, username = ?, password = ? WHERE id = ?"
// 		result, err := urepo.connection.Exec(sqlUpdate, user.Name, user.Email, user.Username, user.Password, user.ID)

// 		affects, err := result.RowsAffected()
// 		lastID, _ := result.LastInsertId()
// 		return urepo.FindByID(int(lastID)), err
// 	}

// 	return user, errors.New("No connection with database")
// }
