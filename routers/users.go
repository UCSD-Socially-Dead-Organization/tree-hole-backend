package routers

import (
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/controllers"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(v1 *gin.RouterGroup) {
	var user controllers.UserController

	v1.GET("/user", user.GetAll)
	v1.POST("/user", user.Create)
	v1.GET("/user/:id", user.GetOne)
	v1.PUT("/user/:id", user.Update)
}
