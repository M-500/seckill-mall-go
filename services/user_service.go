package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"seckill-mall-go/models"
	"seckill-mall-go/repositories"
)

type IUserService interface {
	IsPwdSuccess(userName string, pwd string) (user *models.User, isOk bool)
	AddUser(user *models.User) (uid int64, err error)
}

type UserService struct {
	userRepo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) IUserService {
	return &UserService{
		userRepo: repo,
	}
}

func (u *UserService) IsPwdSuccess(userName string, pwd string) (user *models.User, isOk bool) {
	user, err := u.userRepo.SelectByUsername(userName)
	if err != nil {
		return nil, false
	}
	// 判断密码
	isOk, _ = ValidatePwd(user.Password, pwd)
	if err != nil {
		return &models.User{}, false
	}
	return user, isOk
}

func (u *UserService) AddUser(user *models.User) (uid int64, err error) {
	// 判断用户是否存在
	res, err := u.userRepo.SelectByUsername(user.Username)
	if res != nil {
		return 0, errors.New("用户名已存在")
	}

	// 密码加密
	hashPwd, err := GeneratePwd(user.Password)
	if err != nil {
		return 0, err
	}
	data := &models.User{
		Username: user.Username,
		Password: string(hashPwd),
		NickName: user.NickName,
	}
	return u.userRepo.Insert(data)

}

// 这里密码直接选择哈希
func ValidatePwd(oldPwd string, hasPwd string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(oldPwd), []byte(hasPwd)); err != nil {
		return false, errors.New("密码校验不通过！")
	}
	return true, nil
}

// 生成pwd
func GeneratePwd(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}
