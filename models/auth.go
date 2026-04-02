package models

import "github.com/DrZiMo/watch-tracker-api-golang/db"

type User struct {
	ID        int64
	Name      string `binding:"required"`
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	LastLogin string `binding:"required"`
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
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.LastLogin)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
