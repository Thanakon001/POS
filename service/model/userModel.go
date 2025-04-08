package model

import (
	"database/sql"
)

type User struct {
	ID        int64        `json:"id"`
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	Email     string       `json:"email"`
	Phone     string       `json:"phone"`
	Avatar    string       `json:"avatar"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

func CreateUser(db *sql.DB, user User) error {
	query := "INSERT INTO users (id, name, email) VALUES (?, ?, ?)"
	_, err := db.Exec(query, user.ID, user.Username, user.Email)
	return err
}

func GetUserByID(db *sql.DB, id int64) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password, email, phone, avatar, created_at, updated_at FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Avatar, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return User{}, err
	}
	return user, nil
}
