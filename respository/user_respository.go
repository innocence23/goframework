package respository

import (
	"goframework/lib"
	"goframework/model"
)

type UserRespository struct {
}

func (u *UserRespository) FindById(id int) (*model.User, error) {
	user := &model.User{}
	res := lib.DB.First(&user, id)
	return user, res.Error
}

func (u *UserRespository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	res := lib.DB.Where("email = ?", email).First(&user)
	return user, res.Error
}

func (u *UserRespository) Create(user *model.User) (*model.User, error) {
	res := lib.DB.Debug().Create(user)
	return user, res.Error
}

func (u *UserRespository) Update(user *model.User, data map[string]interface{}) (*model.User, error) {
	res := lib.DB.Debug().Model(user).Updates(data)
	return user, res.Error
}
