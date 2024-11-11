package models

import (
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
