package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigIndex(val gin.H) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "config.html", val)
	}
}
