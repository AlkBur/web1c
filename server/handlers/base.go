package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BaseIndex(val gin.H) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.html", val)
	}
}
