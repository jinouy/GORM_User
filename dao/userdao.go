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

	if err := utils.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserName(u *model.User) (model.User, []byte, error) {

	var users model.User
	var count int64
	err := utils.Db.Model(&users).Where("name = ?", u.Name).Count(&count)
	if err.Error != nil {
		return model.User{}, nil, err.Error
	}
	if count == 0 {
		var data []byte
		data = []byte("名字不存在")
		return model.User{}, data, nil
	}
	err = utils.Db.Where("name = ?", u.Name).First(&users)
	if err.Error != nil {
		return model.User{}, nil, err.Error
	}
	return users, nil, nil
}

func GetUsersPage(p *model.Page) ([]model.User, []byte, error) {

	var users []model.User
	var total int64

	//if err := utils.Db.Model(&model.User{}).Order("created_at DESC").Find(&users).Error; err != nil {
	//	return nil, nil, err
	//}
	//
	if err := utils.Db.Model(&model.User{}).Limit(p.PageSize).Offset((p.PageNum - 1) * p.PageSize).Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, nil, err
	}
	//PageNum := (p.PageNum - 1) * p.PageSize
	//sqlStr := "SELECT * FROM (SELECT * FROM users WHERE users.deleted_at IS NULL LIMIT ? OFFSET ? ) as T  ORDER BY created_at DESC"
	//err := utils.Db.Raw(sqlStr, p.PageSize, PageNum).Scan(&users).Error
	//if err != nil {
	//	return nil, nil, err
	//}

	//if err := utils.Db.Table("(?) as T", utils.Db.Model(&model.User{}).Limit(p.PageSize).Offset((p.PageNum-1)*p.PageSize)).Order("created_at DESC").Find(&users); err != nil {
	//	return nil, nil, nil
	//}
	//if err := utils.Db.Model(&users).Order("created_at DESC").Find(&users).Error; err != nil {
	//	return nil, nil, err
	//}

	if err := utils.Db.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}
	if int64(p.PageSize) > total {
		var data []byte
		data = []byte("查询超出数据长度")
		return nil, data, nil
	}
	return users, nil, nil

}

func AddUser(u *model.User) (model.User, error) {

	user := &model.User{Name: u.Name}
	var users model.User

	err := utils.Db.Create(&user)
	if err.Error != nil {
		return users, err.Error
	}

	err = utils.Db.Where("name = ?", u.Name).First(&users)
	if err.Error != nil {
		return users, err.Error
	}
	return users, nil
}

func DpdUser(u *model.Username) (model.User, error) {

	var users model.User

	err := utils.Db.Model(&users).Where("name = ?", u.OldName).Update("name", u.NewName)
	if err.Error != nil {
		return users, err.Error
	}
	err = utils.Db.Where("name = ?", u.NewName).First(&users)
	if err.Error != nil {
		return users, err.Error
	}
	return users, nil
}

func DelUser(u *model.User) (model.User, error) {

	var users model.User

	err := utils.Db.Clauses(clause.Returning{}).Where("name = ?", u.Name).Delete(&users)
	if err.Error != nil {
		return model.User{}, err.Error
	}
	return users, nil
}

//func DelUserID(u *model.User) ([]model.User, error) {
//
//	//utils.Db.Delete(&model.User{}, u.ID)
//	var users []model.User
//	utils.Db.Clauses(clause.Returning{}).Where("id = ?", u.ID).Delete(&users)
//	return users, nil
//}
