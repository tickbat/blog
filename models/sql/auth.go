package sql

import "blog/models"

func CheckAuth(conditions models.Auth) (models.Auth, error) {
	var auth models.Auth
	err := models.Db.Where(conditions).First(&auth).Error
	return auth, err
}
