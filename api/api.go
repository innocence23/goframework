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
}

func NewServer(db *gorm.DB) *Server {
	server := &Server{
		DB: db,
	}

	server.PostService = &service.PostService{
		PostRespository: &respository.PostRespository{},
	}

	server.initRoutes()
	return server
}

// func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
// 	var err error
// 	if Dbdriver == "mysql" {
// 		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
// 		server.DB, err = gorm.Open(Dbdriver, DBURL)
// 		if err != nil {
// 			fmt.Printf("Cannot connect to %s database", Dbdriver)
// 			log.Fatal("This is the error:", err)
// 		} else {
// 			fmt.Printf("We are connected to the %s database", Dbdriver)
// 		}
// 	}
// 	if Dbdriver == "postgres" {
// 		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
// 		server.DB, err = gorm.Open(Dbdriver, DBURL)
// 		if err != nil {
// 			fmt.Printf("Cannot connect to %s database", Dbdriver)
// 			log.Fatal("This is the error:", err)
// 		} else {
// 			fmt.Printf("We are connected to the %s database", Dbdriver)
// 		}
// 	}
// 	if Dbdriver == "sqlite3" {
// 		//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
// 		server.DB, err = gorm.Open(Dbdriver, DbName)
// 		if err != nil {
// 			fmt.Printf("Cannot connect to %s database\n", Dbdriver)
// 			log.Fatal("This is the error:", err)
// 		} else {
// 			fmt.Printf("We are connected to the %s database\n", Dbdriver)
// 		}
// 		server.DB.Exec("PRAGMA foreign_keys = ON")
// 	}

// 	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

// 	server.Router = mux.NewRouter()

// 	server.initializeRoutes()
// }

func (s *Server) initRoutes() {
	router := gin.Default()
	g := router.Group("v1")
	{
		g.GET("/", s.Home)

		// // Login Route
		// g.POST("/login", s.Login)

		// //Users routes
		// g.POST("/users", s.CreateUser)
		// g.GET("/users", s.GetUsers)
		// g.GET("/users/:id", s.GetUser)
		// g.PUT("/users/:id", s.UpdateUser)
		// g.DELETE("/users/:id", s.DeleteUser)

		//Posts routes
		g.GET("/posts/:id", s.GetPost)
		g.GET("/posts", s.GetPosts)
		g.POST("/posts", s.CreatePost)
		g.PUT("/posts/:id", s.UpdatePost)
		g.DELETE("/posts/:id", s.DeletePost)
	}
	s.Router = router
}

func (s *Server) Run(addr string) {
	fmt.Println(addr)
	log.Fatal(s.Router.Run(addr))
}
