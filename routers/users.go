package routers

import (
	"net/http"
	"time"

	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/repository"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/repository/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UsersRoutes(v1 *gin.RouterGroup, repo repository.UserRepo) {
	userHandler := userHandler{repo: repo}

	v1.GET("/user", userHandler.GetAll)
	v1.POST("/user", userHandler.Create)
	v1.GET("/user/:id", userHandler.GetOne)
	v1.PUT("/user/:id", userHandler.Update)
}

type usersResp struct {
	Users []userResp `json:"users"`
}
type userResp struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"username"`
}

type userHandler struct {
	repo repository.UserRepo
}

func (u *userHandler) GetAll(ctx *gin.Context) {
	var err error
	var results []models.User
	results, err = u.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		ctx.Abort()
		return
	}
	var resps []userResp
	for _, res := range results {
		resp := userResp{
			ID:   res.ID,
			Name: res.Username,
		}
		resps = append(resps, resp)
	}
	ctx.JSON(http.StatusOK, &usersResp{Users: resps})
}

func (u *userHandler) GetOne(ctx *gin.Context) {
	var err error
	var result models.User

	id_str := ctx.Param("id")
	id, err := uuid.Parse(id_str)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid ID"})
		ctx.Abort()
		return
	}
	result, err = u.repo.GetOne(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, &userResp{ID: result.ID, Name: result.Username})
}

type userReq struct {
	Id         uuid.UUID `json:"id"`
	ProfilePic []byte    `json:"profile_pic"`
	Name       string    `json:"username"`
	Age        int       `json:"age"`
	LastLogin  time.Time `json:"last_login"`
}

func (u *userHandler) Create(ctx *gin.Context) {
	var req userReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form", "form": req})
		ctx.Abort()
		return
	}
	user := models.User{
		ID:        req.Id,
		Username:  req.Name,
		Age:       req.Age,
		LastLogin: req.LastLogin,
	}
	u.repo.Create(&user)
	ctx.JSON(http.StatusCreated, &user)
}

func (u *userHandler) Update(ctx *gin.Context) {
	var err error
	var req userReq
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form", "form": req})
		ctx.Abort()
		return
	}
	id_str := ctx.Param("id")
	id, err := uuid.Parse(id_str)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid ID"})
		ctx.Abort()
		return
	}

	user := models.User{
		ID:         req.Id,
		Username:   req.Name,
		Age:        req.Age,
		LastLogin:  req.LastLogin,
		ProfilePic: req.ProfilePic,
	}

	u.repo.Update(id, &user)
	ctx.JSON(http.StatusOK, &user)
}
