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

type Page struct {
	PageNum  int `json:"page_num"  grom:"page_num"`
	PageSize int `json:"page_size" grom:"page_size"`
}

func init() {
	utils.Db.AutoMigrate(&User{})
}
