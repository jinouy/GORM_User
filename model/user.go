package model

import (
	"gorm.io/gorm"
	"gorm_tcp/utils"
)

// User 结构体
type User struct {
	gorm.Model
	Name string `json:"name" grom:"name"`
}
type Username struct {
	OldName string `json:"oldname" grom:"oldname"`
	NewName string `json:"newname" grom:"newname"`
}

func init() {
	utils.Db.AutoMigrate(&User{})
}
