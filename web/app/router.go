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
	// Cookie主要存储在用户的浏览器中，用于在客户端和服务器之间传输和存储信息，而 session则存储在服务器端，
	// 用于在用户访问期间保持用户的状态信息。两者通常一起使用，通过将会话ID存储在Cookie中来实现在多个请求之间跟踪用户的会话状态。

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	// secret key for cookie store
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	// router.GET("/", home.Handler)
	router.GET("/api/auth/login", login.Handler(auth))
	router.GET("/api/auth/callback", callback.Handler(auth))

	router.GET("/api/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	return router, nil
}
