package repositories

import (
	"errors"
	"gorm.io/gorm"
	"seckill-mall-go/common"
	"seckill-mall-go/models"
)

type IUserRepository interface {
	SelectByUsername(username string) (user *models.User, err error)
	SelectByID(uid int64) (user *models.User, err error)
	Insert(user *models.User) (uid int64, err error)
}

type UserManager struct {
	DB *gorm.DB
}

func NewUserRepository() IUserRepository {
	return &UserManager{
		DB: common.DB(),
	}
}

func (u *UserManager) SelectByUsername(username string) (user *models.User, err error) {
	queryData := models.User{}
	exist := u.DB.Where(&models.User{Username: username}).First(&queryData)
	if exist.RowsAffected <= 0 {
		return nil, errors.New("用户不存在")
	}
	return &queryData, err
}

func (u *UserManager) SelectByID(uid int64) (user *models.User, err error) {
	queryData := &models.User{}
	result := u.DB.First(queryData, uid)
	if result.Error != nil {
		return nil, result.Error
	}
	return queryData, nil
}

func (u *UserManager) Insert(user *models.User) (uid int64, err error) {
	queryData := &models.User{}
	exist := u.DB.Where(&models.User{Username: user.Username}).First(&queryData)
	if exist.RowsAffected == 1 {
		return 0, errors.New("用户名已存在")
	}
	createUser := models.User{
		NickName: user.NickName,
		Username: user.Username,
		Password: user.Password,
	}
	u.DB.Create(&createUser)

	return createUser.ID, nil
}
