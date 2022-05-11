package dao

import (
	"gorm.io/gorm/clause"
	"gorm_tcp/model"
	"gorm_tcp/utils"
)

//GetUsers 获取数据库中所有的用户
func GetUsers() ([]model.User, error) {
	//执行sql操作
	var users []model.User

	err := utils.Db.Find(&users)
	if err.Error != nil {
		return nil, err.Error
	}

	return users, nil
}

func GetUserName(u *model.User) (model.User, []byte, error) {

	var users model.User
	var count int64
	err := utils.Db.Model(&users).Where("name = ?", u.Name).Count(&count)
	if err.Error != nil {
		return users, nil, err.Error
	}
	if count == 0 {
		var data []byte
		data = []byte("名字不存在")
		return users, data, nil
	}
	err = utils.Db.Where("name = ?", u.Name).First(&users)
	if err.Error != nil {
		return users, nil, err.Error
	}
	return users, nil, nil
}

func AddUser(u *model.User) (model.User, []byte, error) {
	user := &model.User{Name: u.Name}
	var users model.User
	var count int64
	err := utils.Db.Model(&user).Where("name = ?", u.Name).Count(&count)
	if err.Error != nil {
		return users, nil, err.Error
	}
	if count > 0 {
		var data []byte
		data = []byte("这个名字已经存在")
		return users, data, nil
	}
	err = utils.Db.Create(&user)
	if err.Error != nil {
		return users, nil, err.Error
	}

	err = utils.Db.Where("name = ?", u.Name).First(&users)
	if err.Error != nil {
		return users, nil, err.Error
	}
	return users, nil, nil
}

func DpdUser(u *model.Username) (model.User, []byte, error) {

	var users model.User
	var count int64
	err := utils.Db.Model(&users).Where("name = ?", u.OldName).Count(&count)
	if err.Error != nil {
		return users, nil, err.Error
	}
	if count == 0 {
		var data []byte
		data = []byte("需要修改的名字不存在")
		return users, data, nil
	}
	err = utils.Db.Model(&users).Where("name = ?", u.NewName).Count(&count)
	if err.Error != nil {
		return users, nil, err.Error
	}
	if count > 0 {
		var data []byte
		data = []byte("名字已经存在")
		return users, data, nil
	}

	err = utils.Db.Model(&users).Where("name = ?", u.OldName).Update("name", u.NewName)
	if err.Error != nil {
		return users, nil, err.Error
	}
	err = utils.Db.Where("name = ?", u.NewName).First(&users)
	if err.Error != nil {
		return users, nil, err.Error
	}
	return users, nil, nil
}

func DelUser(u *model.User) (model.User, []byte, error) {

	var users model.User
	var count int64
	err := utils.Db.Model(&users).Where("name = ?", u.Name).Count(&count)
	if err.Error != nil {
		return users, nil, err.Error
	}
	if count == 0 {
		var data []byte
		data = []byte("需要删除的名字不存在")
		return users, data, nil
	}
	err = utils.Db.Clauses(clause.Returning{}).Where("name = ?", u.Name).Delete(&users)
	if err.Error != nil {
		return users, nil, err.Error
	}
	return users, nil, nil
}

//func DelUserID(u *model.User) ([]model.User, error) {
//
//	//utils.Db.Delete(&model.User{}, u.ID)
//	var users []model.User
//	utils.Db.Clauses(clause.Returning{}).Where("id = ?", u.ID).Delete(&users)
//	return users, nil
//}
