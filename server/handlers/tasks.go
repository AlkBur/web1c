package handlers

import (
	"github.com/AlkBur/web1c/server/db1c"
	"github.com/AlkBur/web1c/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTasks(db *db1c.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PostTasks(db *db1c.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.PostTasks(db))
	}
}
