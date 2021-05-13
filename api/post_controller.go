package api

import (
	"goframework/lib"
	"goframework/model"

	"github.com/gin-gonic/gin"
)

type PostIdParam struct {
	Id *int `uri:"id" binding:"required,gt=0" label:"ID"`
}

func (s *Server) GetPost(c *gin.Context) (data interface{}, error *lib.Error) {
	var param PostIdParam
	if err := c.ShouldBindUri(&param); err != nil {
		error = lib.NewAutoParamError(err)
		return
	}
	data, err := s.PostService.GetById(*param.Id)
	if err != nil {
		error = lib.NewNotFoundError(err.Error())
		return
	}
	return
}

type GetPostsParam struct {
	Offset *int `form:"offset" binding:"required,gte=0" label:"偏移量"` //使用指针可以解决required 传0的情况
	Limit  int  `form:"limit" binding:"required,gt=0" label:"每页个数"`
}

func (s *Server) GetPosts(c *gin.Context) (data interface{}, error *lib.Error) {
	var param GetPostsParam
	if err := c.ShouldBind(&param); err != nil {
		error = lib.NewAutoParamError(err)

		return
	}
	data, err := s.PostService.GetByPage(param.Limit, *param.Offset)
	if err != nil {
		error = lib.NewNotFoundError(err.Error())
		return
	}
	return
}

type CreatePostParam struct {
	Title   string `json:"title" binding:"required" label:"标题"`
	Desc    string `json:"desc" binding:"required" label:"描述"`
	Content string `json:"content" binding:"required" label:"内容"`
}

func (s *Server) CreatePost(c *gin.Context) (data interface{}, error *lib.Error) {
	var param CreatePostParam
	if err := c.ShouldBind(&param); err != nil {
		error = lib.NewAutoParamError(err)
		return
	}
	post := &model.Post{
		Title:   param.Title,
		Desc:    param.Desc,
		Content: param.Content,
		Status:  1,
	}
	data, err := s.PostService.Create(post)
	if err != nil {
		error = lib.NewInternalError(err.Error())
		return
	}
	return
}

type UpdatePostParam struct {
	Title   string `json:"title" binding:"required" label:"标题"`
	Desc    string `json:"desc" binding:"required" label:"描述"`
	Content string `json:"content" binding:"required" label:"内容"`
}

func (s *Server) UpdatePost(c *gin.Context) (data interface{}, error *lib.Error) {
	var paramId PostIdParam
	if err := c.ShouldBindUri(&paramId); err != nil {
		error = lib.NewAutoParamError(err)
		return
	}
	var param UpdatePostParam
	if err := c.ShouldBind(&param); err != nil {
		error = lib.NewAutoParamError(err)
		return
	}
	updateData := map[string]interface{}{
		"title":   param.Title,
		"desc":    param.Desc,
		"content": param.Content,
	}
	data, err := s.PostService.UpdateById(*paramId.Id, updateData)
	if err != nil {
		error = lib.NewInternalError(err.Error())
		return
	}
	return
}

func (s *Server) DeletePost(c *gin.Context) (data interface{}, error *lib.Error) {
	var paramId PostIdParam
	if err := c.ShouldBindUri(&paramId); err != nil {
		error = lib.NewAutoParamError(err)
		return
	}
	if err := s.PostService.DeleteById(*paramId.Id); err != nil {
		error = lib.NewNotFoundError(err.Error())
		return
	}
	return
}
