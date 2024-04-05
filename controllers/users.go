package controllers

import (
	"net/http"

	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/database"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/logger"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctrl *UserController) CreateUser(ctx *gin.Context) {
	user := new(models.User)

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DB.Create(&user).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &user)
}

func (ctrl *UserController) GetUsers(ctx *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{"data": users})

}
