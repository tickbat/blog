package models

type Auth struct {
	ID       *int    `gorm:"primary_key" json:"id"`
	Username *string `json:"username" binding:"required"`
	Password *string `json:"password" binding:"required"`
}

func CheckAuth(conditions Auth) bool {
	var auth Auth
	Db.Where(conditions).First(&auth)
	if auth.ID == nil && auth.Password != conditions.Password {
		return false
	}
	return true
}
