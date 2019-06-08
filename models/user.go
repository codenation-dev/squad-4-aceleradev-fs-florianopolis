package models

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/utils"
	// "codenation/squad-4-aceleradev-fs-florianopolis/models"

	"fmt"
)

// User of the app
type User struct {
	UID   int32  `json:"uid"`
	Login string `json:"login"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

// NewUser insert a new user on the db
func NewUser(u User) (bool, error) {
	db := Connect()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	// TODO: quando tenta inserir com valor duplo acaba consumindo uma uid e
	// pulando um numero, tem como não acontecer isso?
	sql := `INSERT INTO users (login, email, pass)
			VALUES ($1, $2, $3)
			RETURNING uid`
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()

		hashedPassword, err := utils.Bcrypt(u.Pass)
		if err != nil {
			tx.Rollback()
			return false, err
		}

		err = stmt.QueryRow(
			u.Login, u.Email, string(hashedPassword),
		).Scan(&u.UID)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	return true, tx.Commit()
}

// GetUsers returns a list of all users of the app
func GetUsers() ([]User, error) {
	db := Connect()
	defer db.Close()
	sql := "SELECT uid, login, email FROM users"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.UID, &u.Login, &u.Email)
		u.Pass = "*****"
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
