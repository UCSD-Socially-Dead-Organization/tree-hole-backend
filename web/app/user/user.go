package user

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	log.Println("user successfully authenticated")
	log.Println(profile)

	// TODO: Might need to generate a user profile in the database if not exists
	ctx.Redirect(http.StatusTemporaryRedirect, "/user")
}

type Bottles struct {
	Bottles []Bottle `json:"bottles"`
}

type Bottle struct {
}

func GetBottlesHandler(ctx *gin.Context) {
	// session := sessions.Default(ctx)
	// profile := session.Get("profile")

	// // TODO: get bottles from database

	// ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func PutBottlesHandler(ctx *gin.Context) {
	// session := sessions.Default(ctx)
	// profile := session.Get("profile")
}
