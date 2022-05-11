package dao

import (
	"errors"
	"gorm.io/gorm/clause"
	"gorm_tcp/model"
	"gorm_tcp/utils"
)

//GetUsers 获取数据库中所有的用户
func GetUsers() ([]model.User, error) {
	//执行sql操作
	var users []model.User

	utils.Db.Find(&users)

	return users, nil
}

func GetUserName(u *model.User) ([]model.User, error) {

	var users []model.User
	var count int64
	utils.Db.Model(&users).Where("name = ?", u.Name).Count(&count)
	if count == 0 {
		return nil, errors.New("名字不存在")
	}
	utils.Db.Where("name = ?", u.Name).First(&users)

	return users, nil
}

func AddUser(u *model.User) ([]model.User, error) {
	user := &model.User{Name: u.Name}

	var count int64
	utils.Db.Model(&user).Where("name = ?", u.Name).Count(&count)
	if count > 0 {
		return nil, errors.New("这个名字已经存在")
	}
	utils.Db.Create(&user)

	var users []model.User
	utils.Db.Where("name = ?", u.Name).First(&users)
	return users, nil
}

func DpdUser(u *model.Username) ([]model.User, error) {

	var users []model.User
	var count int64
	utils.Db.Model(&users).Where("name = ?", u.OldName).Count(&count)
	if count == 0 {
		return nil, errors.New("需要修改的名字不存在")
	}
	utils.Db.Model(&users).Where("name = ?", u.NewName).Count(&count)
	if count > 0 {
		return nil, errors.New("名字已经存在")
	}

	utils.Db.Model(&users).Where("name = ?", u.OldName).Update("name", u.NewName)
	utils.Db.Where("name = ?", u.NewName).First(&users)
	return users, nil
}

func DelUser(u *model.User) ([]model.User, error) {

	var users []model.User
	var count int64
	utils.Db.Model(&users).Where("name = ?", u.Name).Count(&count)
	if count == 0 {
		return nil, errors.New("需要删除的名字不存在")
	}
	utils.Db.Clauses(clause.Returning{}).Where("name = ?", u.Name).Delete(&users)
	return users, nil
}

//func DelUserID(u *model.User) ([]model.User, error) {
//
//	//utils.Db.Delete(&model.User{}, u.ID)
//	var users []model.User
//	utils.Db.Clauses(clause.Returning{}).Where("id = ?", u.ID).Delete(&users)
//	return users, nil
//}
