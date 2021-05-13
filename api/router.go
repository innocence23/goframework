package api

import (
	"goframework/api/middleware"

	"github.com/gin-gonic/gin"
)

func (s *Server) initRoutes() {
	router := gin.Default()
	g := router.Group("v1")
	{
		g.GET("/", s.Home)

		//注册登陆
		g.POST("/register", errorWrapper(s.Register))
		g.POST("/login", errorWrapper(s.Login))
		g.POST("/logout", errorWrapper(s.Logout))
		gw := g.Use(middleware.JWTAuth())
		{
			gw.GET("/profile", errorWrapper(s.Profile))

			//Posts routes
			gw.GET("/posts/:id", errorWrapper(s.GetPost))
			gw.GET("/posts", errorWrapper(s.GetPosts))
			gw.POST("/posts", errorWrapper(s.CreatePost))
			gw.PUT("/posts/:id", errorWrapper(s.UpdatePost))
			gw.DELETE("/posts/:id", errorWrapper(s.DeletePost))
		}

	}
	s.Router = router
}
