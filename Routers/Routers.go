package Routers

import (
	"../Controllers"
	_ "../docs" //fix error not yet registered swag
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
