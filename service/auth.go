package service

import (
	"blog/models"
	"blog/models/sql"
)

// 0 验证通过，1 用户不存在， 2 密码或者用户名错误

func CheckAuth(auth models.Auth) (int, error) {
	result, err := sql.CheckAuth(auth.Username)
	if err != nil {
		return 0, err
	}
	if result.ID == 0 {
		return 1, nil
	}
	if auth.Password != result.Password {
		return 2, nil
	}
	return 0, nil
}
