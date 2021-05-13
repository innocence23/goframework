package api

import (
	"goframework/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WrapperHandle func(c *gin.Context) (interface{}, *lib.Error)

func ErrorWrapper(handle WrapperHandle) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handle(c)
		if err != nil {
			c.JSON(http.StatusOK, lib.ErrorResponse(err))
		} else {
			c.JSON(http.StatusOK, lib.SuccessResponse(data))
		}
	}
}
