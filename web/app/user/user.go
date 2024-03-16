package user

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler for our logged-in user page.
func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	ctx.JSON(http.StatusOK, gin.H{
		"profile": profile,
	})
	// TODO: Add a template for this.
	fmt.Println("FUCK YOU!")
}
