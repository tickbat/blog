package service

import (
	"blog/models"
	"blog/models/sql"
)

func CheckAuth(conditions models.Auth) (bool, error) {
	auth, err := sql.CheckAuth(conditions)
	if err != nil {
		return false, err
	}
	if auth.ID == 0 || auth.Password != conditions.Password {
		return false, nil
	}
	return true, nil
}
