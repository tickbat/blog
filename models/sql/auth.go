package sql

import (
	"blog/models"
	"github.com/jinzhu/gorm"
)

func CheckAuth(id string) (models.Auth, error) {
	var auth models.Auth
	if err := models.Db.Where("username = ?", id).First(&auth).Error; err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}
	return auth, nil
}
