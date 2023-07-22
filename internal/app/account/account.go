package account

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username string `json:"username" column:"username"`
	Password string `json:"password" column:"password"`
}

func (u *Account) TableName() string {
	return "accounts"
}
