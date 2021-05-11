package respository

import (
	"goframework/lib"
	"goframework/model"
)

type PostRespository struct {
}

func (p *PostRespository) GetById(id int) (*model.Post, error) {
	post := &model.Post{}
	res := lib.DB.First(&post, id)
	return post, res.Error
}
