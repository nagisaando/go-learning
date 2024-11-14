package models

import (
	"errors"

	"example.com/rest-api-2/db"
	"example.com/rest-api-2/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user User) Save() error {

	query := `
	INSERT INTO users (email, password)
	VALUES(?, ?)
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	user.ID = id

	return err

}

func (user *User) ValidateCredential() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string

	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.IsValidPassword(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil

}
