package Routers

import (
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	root := r.Group("/api")
	{
		root.GET("/user", Controllers.GetAllUsers)
		root.GET("/user/:id", Controllers.GetUserById)
		root.POST("/user", Controllers.AddUser)
		root.PUT("/user", Controllers.UpdateUser)
		root.DELETE("/user/:id", Controllers.DeleteUser)
	}
	return r
}
