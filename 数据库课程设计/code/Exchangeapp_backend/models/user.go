package models

import (
	"exchangeapp/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	// 建立一对多关系
	Articles []Article `json:"articles" gorm:"foreignKey:UserID"`
}

// BeforeSave 在保存前自动加密密码
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// 只有当密码字段被修改或者是新创建时才加密
	if len(u.Password) > 0 {
		hash, err := utils.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hash
	}
	return
}
