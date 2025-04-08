package controller

import (
	"database/sql"
	"pos/service/model"
)

type UserController struct {
	DB *sql.DB
}

func (uc *UserController) CreateUser(user model.User) error {
	return model.CreateUser(uc.DB, user)
}
