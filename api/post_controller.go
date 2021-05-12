package api

import (
	"fmt"
	"goframework/model"
	"goframework/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostIdParam struct {
	Id int `uri:"id" binding:"required,gt=0" label:"ID"`
}

func (s *Server) GetPost(c *gin.Context) {
	var param PostIdParam
	if err := c.ShouldBindUri(&param); err != nil {
		msg := util.ValidateParams(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	result, err := s.PostService.GetById(param.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

type GetPostsParam struct {
	Offset *int `form:"offset" binding:"required,gte=0" label:"偏移量"` //使用指针可以解决required 传0的情况
	Limit  int  `form:"limit" binding:"required,gt=0" label:"每页个数"`
}

func (s *Server) GetPosts(c *gin.Context) {
	var param GetPostsParam
	if err := c.ShouldBind(&param); err != nil {
		msg := util.ValidateParams(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	result, err := s.PostService.GetByPage(param.Limit, *param.Offset)
	fmt.Println(result, err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

type CreatePostParam struct {
	Title   string `json:"title" binding:"required" label:"标题"`
	Desc    string `json:"desc" binding:"required" label:"描述"`
	Content string `json:"content" binding:"required" label:"内容"`
}

func (s *Server) CreatePost(c *gin.Context) {
	var param CreatePostParam
	if err := c.ShouldBind(&param); err != nil {
		msg := util.ValidateParams(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	post := &model.Post{
		Title:   param.Title,
		Desc:    param.Desc,
		Content: param.Content,
		Status:  1,
	}
	result, err := s.PostService.Create(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

type UpdatePostParam struct {
	Title   string `json:"title" binding:"required" label:"标题"`
	Desc    string `json:"desc" binding:"required" label:"描述"`
	Content string `json:"content" binding:"required" label:"内容"`
}

func (s *Server) UpdatePost(c *gin.Context) {
	var paramId PostIdParam
	if err := c.ShouldBindUri(&paramId); err != nil {
		msg := util.ValidateParams(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	var param UpdatePostParam
	if err := c.ShouldBind(&param); err != nil {
		msg := util.ValidateParams(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	updateData := map[string]interface{}{
		"title":   param.Title,
		"desc":    param.Desc,
		"content": param.Content,
	}
	result, err := s.PostService.UpdateById(paramId.Id, updateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (s *Server) DeletePost(c *gin.Context) {
	var paramId PostIdParam
	if err := c.ShouldBindUri(&paramId); err != nil {
		msg := util.ValidateParams(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	if err := s.PostService.DeleteById(paramId.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ""})
}
