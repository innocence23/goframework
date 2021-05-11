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

func (p *PostRespository) GetByPage(limit, offset int) ([]model.Post, error) {
	posts := make([]model.Post, 0)
	res := lib.DB.Limit(limit).Offset(offset).Find(&posts)
	return posts, res.Error
}

func (p *PostRespository) DeleteById(id int) error {
	res := lib.DB.Delete(&model.Post{}, id)
	return res.Error
}

func (p *PostRespository) Update(post *model.Post, data map[string]interface{}) (*model.Post, error) {
	res := lib.DB.Debug().Model(post).Updates(data)
	return post, res.Error
}

func (p *PostRespository) Create(post *model.Post) (*model.Post, error) {
	res := lib.DB.Debug().Create(post)
	return post, res.Error
}
