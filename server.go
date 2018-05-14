package main

import (
	// Standard library packages

	// Third party packages

	"mugg/mongo-rest/conf"
	"mugg/mongo-rest/controllers"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})

	// // Output to stdout instead of the default stderr
	// // Can be any io.Writer, see below for File example
	// log.SetOutput(os.Stdout)

	// // Only log the warning severity or above.
	// log.SetLevel(log.WarnLevel)

}

func main() {

	config := conf.Config

	// Get a UserController instance
	uc := controllers.NewUserController(mogoSession())
	lg := controllers.NewLogController(mogoSession())

	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// Get a user resource
	router := gin.Default()

	api := router.Group("/api")
	// Simple group: v1
	v1 := api.Group("/v1")
	{
		v1.GET("/users", uc.UsersList)
		v1.GET("/users/:id", uc.GetUser)
		v1.POST("/users", uc.CreateUser)
		v1.PUT("/users/:id", uc.UpdateUser)
		v1.DELETE("/users/:id", uc.RemoveUser)

		v1.POST("/logs", lg.Create)

	}

	router.Run(":" + config.APP.Port)
}

// mogoSession creates a new mongo session and panics if connection error occurs
func mogoSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://test:test@118.25.45.204:27018/taizhou")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
