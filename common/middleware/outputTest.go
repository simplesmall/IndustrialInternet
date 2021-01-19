package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleOutTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"OKK",
		})
	}
}