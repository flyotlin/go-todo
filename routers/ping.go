package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}
