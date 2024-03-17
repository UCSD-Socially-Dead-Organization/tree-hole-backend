package restful

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/treehole-backend/envconfig"
	"github.com/treehole-backend/platform/authenticator"
	"github.com/treehole-backend/platform/middleware"
	"github.com/treehole-backend/web/app/callback"
	"github.com/treehole-backend/web/app/login"
	"github.com/treehole-backend/web/app/logout"
	"github.com/treehole-backend/web/app/user"
)

func Register(env envconfig.Env, auth *authenticator.Authenticator) (*gin.Engine, error) {
	router := gin.Default()
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	// secret key for cookie store
	store := cookie.NewStore([]byte("secret"))
	// 存了個 auth-session 的 cookie
	router.Use(sessions.Sessions("auth-session", store))

	router.GET("/api/auth/login", login.Handler(auth))
	router.GET("/api/auth/callback", callback.Handler(auth))

	router.GET("/api/user", middleware.IsAuthenticated, user.Handler)

	// Bottle services
	router.GET("/api/user/bottles", middleware.IsAuthenticated, user.GetBottlesHandler)
	router.PUT("/api/user/bottles", middleware.IsAuthenticated, user.PutBottlesHandler)
	router.GET("/logout", logout.Handler)

	return router, nil
}
