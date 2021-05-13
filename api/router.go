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

		//注册登陆
		g.POST("/register",  ErrorWrapper(s.Register))
		g.POST("/login",  ErrorWrapper(s.Login))
		g.POST("/logout",  ErrorWrapper(s.Logout))

		g.GET("/profile", ErrorWrapper(s.Profile))

		//Posts routes
		g.GET("/posts/:id", ErrorWrapper(s.GetPost))
		g.GET("/posts", ErrorWrapper(s.GetPosts))
		g.POST("/posts", ErrorWrapper(s.CreatePost))
		g.PUT("/posts/:id", ErrorWrapper(s.UpdatePost))
		g.DELETE("/posts/:id", ErrorWrapper(s.DeletePost))
	}
	s.Router = router
}
