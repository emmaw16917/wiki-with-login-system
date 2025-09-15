package repository

import (
	"database/sql"
	"wiki/backend/internal/model"
)

func CreateUser(db *sql.DB, username, passwordHash string) (int64, error) {
	result, err := db.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", username, passwordHash)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetUserByUsername(db *sql.DB, username string) (*model.User, error) {
	user := &model.User{}
	err := db.QueryRow("SELECT id, username, password_hash FROM users WHERE username = ?", username).
		Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}
