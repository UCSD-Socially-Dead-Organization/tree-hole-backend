package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/repository"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/repository/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MatchesRoutes(v1 *gin.RouterGroup, repo repository.MatchRepo) {
	matchHandler := matchHandler{repo: repo}

	v1.GET("/match", matchHandler.GetAll)
	v1.POST("/match", matchHandler.Create)
	v1.GET("/match/:id", matchHandler.GetOne)
	v1.PUT("/match/:id", matchHandler.Update)
}

type matchesResp struct {
	Matches []matchResp `json:"matches"`
}

type matchResp struct {
	ID        uuid.UUID `json:"id"`
	User1     string    `json:"user1"`
	User2     string    `json:"user2"`
	CreatedAt time.Time `json:"created_at"`
}

type matchCreate struct {
	User1 string `json:"user1" binding:"required"`
	User2 string `json:"user2" binding:"required"`
}

type matchUpdate struct {
	Id    uuid.UUID `json:"id" binding:"required"`
	User1 string    `json:"user1"`
	User2 string    `json:"user2"`
}

type matchHandler struct {
	repo repository.MatchRepo
}

func (u *matchHandler) GetAll(ctx *gin.Context) {
	var err error
	var results []models.Match
	results, err = u.repo.GetAll()
	fmt.Println(results)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		ctx.Abort()
		return
	}
	var resps []matchResp
	for _, res := range results {
		resp := matchResp{
			ID:        res.ID,
			User1:     res.User1,
			User2:     res.User2,
			CreatedAt: res.CreatedAt,
		}
		resps = append(resps, resp)
	}
	ctx.JSON(http.StatusOK, &matchesResp{Matches: resps})
}

func (u *matchHandler) GetOne(ctx *gin.Context) {
	var err error
	var result models.Match

	id_str := ctx.Param("id")
	id, err := uuid.Parse(id_str)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid ID"})
		ctx.Abort()
		return
	}
	result, err = u.repo.GetOne(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "match not found"})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, &matchResp{
		ID:    result.ID,
		User1: result.User1,
		User2: result.User2,
	})
}

func (u *matchHandler) Create(ctx *gin.Context) {
	var req matchCreate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form", "error": err.Error()})
		ctx.Abort()
		return
	}

	match := models.Match{
		User1: req.User1,
		User2: req.User2,
	}
	u.repo.Create(&match)

	ctx.JSON(http.StatusCreated, &match)
}

func (u *matchHandler) Update(ctx *gin.Context) {
	var req matchUpdate
	if err := ctx.ShouldBindJSON(&req); err != nil {
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
	if id != req.Id {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "ID in URL does not match ID in form"})
		ctx.Abort()
		return
	}

	match := models.Match{
		ID: req.Id,
	}

	u.repo.Update(&match)
	ctx.JSON(http.StatusOK, &match)
}
