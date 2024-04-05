package controllers

import (
	"net/http"
	"strconv"

	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/database"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/logger"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctrl *UserController) Create(ctx *gin.Context) {
	user := new(models.User)

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form", "form": user})
		ctx.Abort()
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

func (ctrl *UserController) GetAll(ctx *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{"data": users})

}

func (ctrl *UserController) GetOne(ctx *gin.Context) {
	var user models.User

	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.First(&user, userId)
	ctx.JSON(http.StatusOK, gin.H{"data": user})

}

func (ctrl *UserController) Update(ctx *gin.Context) {
	// TODO
	var user models.User
	database.DB.Find(&user)
	ctx.JSON(http.StatusOK, gin.H{"data": user})

}
