package api

import (
	"fmt"
	"goframework/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type GetPostParam struct {
	Id int `uri:"id"`
}

func (s *Server) GetPost(c *gin.Context) {
	var param GetPostParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

func (s *Server) GetPosts(c *gin.Context) {
	var param GetPostsParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := s.PostService.GetByPage(param.Limit, param.Offset)
	fmt.Println(result, err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

type CreatePostParam struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func (s *Server) CreatePost(c *gin.Context) {
	var param CreatePostParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func (s *Server) UpdatePost(c *gin.Context) {
	var param UpdatePostParam
	urlId := c.Param("id")
	id := cast.ToInt(urlId)
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url参数错误"})
		return
	}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateData := map[string]interface{}{
		"title":   param.Title,
		"desc":    param.Desc,
		"content": param.Content,
	}
	result, err := s.PostService.UpdateById(id, updateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

type DeletePostParam struct {
	Id int `uri:"id"`
}

func (s *Server) DeletePost(c *gin.Context) {
	var param DeletePostParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.PostService.DeleteById(param.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ""})
}
