package service

import (
	"database/sql"
	"errors"
	"wiki/backend/internal/model"
	"wiki/backend/internal/repository"
	"wiki/backend/pkg/utils"
)

func RegisterUser(db *sql.DB, username, password string) (int64, error) {
	user, _ := repository.GetUserByUsername(db, username)
	if user != nil {
		return 0, errors.New("用户名已存在")
	}
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return 0, err
	}
	return repository.CreateUser(db, username, passwordHash)
}
func LoginUser(db *sql.DB, username, password string) (*model.User, error) {
	user, err := repository.GetUserByUsername(db, username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if !utils.CheckPassword(password, user.PasswordHash) {
		return nil, errors.New("密码错误")
	}
	return user, nil
}
func GenerateToken(userID int) (string, error) {
	return utils.GenerateJWT(userID)
}
