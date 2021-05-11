package service

import (
	"goframework/model"
	"goframework/respository"
)

type PostService struct {
	PostRespository *respository.PostRespository
}

func (p *PostService) GetById(id int) (*model.Post, error) {
	return p.PostRespository.GetById(id)
}
