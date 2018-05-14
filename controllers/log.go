package controllers

import (
	"mugg/mongo-rest/models"
	"mugg/mongo-rest/tools"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

type (
	// LogController represents the controller for operating on the log resource
	LogController struct {
		base    BaseController
		session *mgo.Session
	}
)

// NewLogController is
func NewLogController(s *mgo.Session) *LogController {
	return &LogController{session: s}
}

// Create creates a new Log resource
func (b LogController) Create(c *gin.Context) {

	var json models.Log

	c.Bind(&json)

	// Write the user to mongo
	// err := b.session.DB(DB_NAME).C("logs").Insert(&json)
	// checkErrTypeOne(err, "Insert failed", "403", c)

	tools.Log.Info(json.Info)
	b.base.JSONSuccess(c, json)
}
