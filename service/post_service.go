package service

import (
	"goframework/model"
	"goframework/respository"
)

type PostService struct {
	PostRespository *respository.PostRespository
}

func (p *PostService) FindById(id int) (*model.Post, error) {
	return p.PostRespository.FindById(id)
}

func (p *PostService) FindByPage(limit, offset int) ([]model.Post, error) {
	return p.PostRespository.FindByPage(limit, offset)
}

func (p *PostService) DeleteById(id int) error {
	if _, err := p.PostRespository.FindById(id); err != nil {
		return err
	}
	return p.PostRespository.DeleteById(id)
}

func (p *PostService) UpdateById(id int, data map[string]interface{}) (*model.Post, error) {
	post, err := p.PostRespository.FindById(id)
	if err != nil {
		return nil, err
	}
	return p.PostRespository.Update(post, data)
}

func (p *PostService) Create(post *model.Post) (*model.Post, error) {
	return p.PostRespository.Create(post)
}
