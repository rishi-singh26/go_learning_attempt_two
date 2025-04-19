package models

import (
	"errors"

	"rishisingh.in/event-manager/db"
	"rishisingh.in/event-manager/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users (email, password) values (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashPass)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&retrivedPassword)

	if err != nil {
		return err
	}

	isValidPass := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !isValidPass {
		return errors.New("invalid credentials")
	}

	return nil
}
