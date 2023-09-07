package db

import (
	"booking-app/pkg/models"
)

func CreateUser(user models.User) error {
	db, err := InitDb()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (username, email, password, github) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password, user.Github)
	if err != nil {
		return err
	}

	return nil
}