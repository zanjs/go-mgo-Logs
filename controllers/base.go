package controllers

import (
	"github.com/gin-gonic/gin"
)

type (
	// BaseController is
	BaseController struct{}
)

// JSONSuccess is res success
func (b BaseController) JSONSuccess(c *gin.Context, json interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, json)
}
