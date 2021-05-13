package service

import (
	"errors"
	"goframework/model"
	"goframework/respository"
	"goframework/util"
)

type UserService struct {
	UserRespository *respository.UserRespository
}

func (u *UserService) FindById(id int) (*model.User, error) {
	return u.UserRespository.FindById(id)
}

func (u *UserService) Register(user *model.User) (*model.User, error) {
	user.Password, _ = util.PasswordHash(user.Password)
	return u.UserRespository.Create(user)
}

func (u *UserService) Login(email, pwd string) (*model.User, error) {
	user, err := u.UserRespository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if flag := util.PasswordVerify(pwd, user.Password); !flag {
		err := errors.New("账号或密码错误")
		return nil, err
	}
	return user, nil
}

func (u *UserService) Logout(user *model.User) (*model.User, error) {
	return u.UserRespository.Create(user)
}

func (u *UserService) UpdateById(id int, data map[string]interface{}) (*model.User, error) {
	user, err := u.UserRespository.FindById(id)
	if err != nil {
		return nil, err
	}
	return u.UserRespository.Update(user, data)
}
