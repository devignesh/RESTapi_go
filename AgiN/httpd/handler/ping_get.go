package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingget(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"Name": "vicky",
	})
}
