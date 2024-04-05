package routers

import (
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/controllers"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(route *gin.Engine) {
	var ctrl controllers.UserController
	v1 := route.Group("/v1/")
	v1.GET("user/", ctrl.GetUsers)
	v1.POST("user/", ctrl.CreateUser)
}
