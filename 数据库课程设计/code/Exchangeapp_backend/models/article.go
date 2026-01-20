package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `binding:"required"`
	Content string `binding:"required"`
	Preview string `binding:"required"`
	// 新增外键关联
	UserID uint `json:"userId"`
	User   User `json:"-" gorm:"foreignKey:UserID"` // 关联 User 模型
}
