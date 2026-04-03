package models

import (
	"errors"
	"time"

	"github.com/DrZiMo/watch-tracker-api-golang/db"
	"github.com/DrZiMo/watch-tracker-api-golang/utils"
)

type User struct {
	ID        int64
	Name      string
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	LastLogin string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetUsers() ([]User, error) {
	query := `SELECT * FROM users`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.LastLogin, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *User) Login() error {
	query := `SELECT id, name, password, last_login FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPass string
	err := row.Scan(&u.ID, &u.Name, &retrivedPass, &u.LastLogin)

	if err != nil {
		return err
	}

	validPassword := utils.CheckPassword(u.Password, retrivedPass)

	if !validPassword {
		return errors.New("Credentials invalid")
	}

	return nil
}

func (u *User) Create() error {
	query := `
	INSERT INTO users (name, email, password, last_login)
	VALUES (?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	timeNow := time.Now().Format(time.RFC3339)
	result, err := stmt.Exec(u.Name, u.Email, hashedPass, timeNow)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = userId
	u.Password = hashedPass
	u.LastLogin = timeNow

	return nil
}
