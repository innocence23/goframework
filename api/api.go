package api

import (
	"fmt"
	"goframework/respository"
	"goframework/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	DB          *gorm.DB
	Router      *gin.Engine
	PostService *service.PostService
	UserService *service.UserService
}

func NewServer(db *gorm.DB) *Server {
	server := &Server{
		DB: db,
	}
	server.PostService = &service.PostService{
		PostRespository: &respository.PostRespository{},
	}
	server.UserService = &service.UserService{
		UserRespository: &respository.UserRespository{},
	}
	server.initRoutes()
	return server
}

func (s *Server) Run(addr string) {
	fmt.Println(addr)
	log.Fatal(s.Router.Run(addr))
}
