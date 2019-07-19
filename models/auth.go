package models

type Auth struct {
	ID *int `gorm:"primary_key" json:"id"`
	Username *string `json:"username" binding:"required"`
	Password *string `json:"password" binding:"required"`
}

func CheckAuth(auth Auth) bool {
	db.Select("id").Where(auth).First(&auth)
	if auth.ID == nil {
		return false
	}
	return true
}