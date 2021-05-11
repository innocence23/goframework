package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome To This Awesome API")
}
