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
		g.GET("/posts/:id", ErrorWrapper(s.GetPost))
		g.GET("/posts", ErrorWrapper(s.GetPosts))
		g.POST("/posts", ErrorWrapper(s.CreatePost))
		g.PUT("/posts/:id", ErrorWrapper(s.UpdatePost))
		g.DELETE("/posts/:id", ErrorWrapper(s.DeletePost))
	}
	s.Router = router
}
