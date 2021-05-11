package api

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) initRoutes() {
	router := gin.Default()
	g := router.Group("v1")
	{
		g.GET("/", s.Home)

		// Login Route
		// g.POST("/login", s.Login)

		//Users routes
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
