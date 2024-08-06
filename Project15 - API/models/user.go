package models

import (
	"api_example/db"
	"api_example/utils"
	"errors"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users (email, password) 
	VALUES (?,?);`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = id

	return nil
}

func (u *User) Authenticate() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.ComparePasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
