package routers

import (
	"net/http"

	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/auth0"
	config "github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/config"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/database"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/repository"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/routers/middleware"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

func Register(gorm *database.GormDatabase, conf *config.Configuration) *gin.Engine {

	is_debug := conf.Server.Debug
	if is_debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Set up the router
	allowedHosts := conf.Server.Allowed_Hosts
	router := gin.New()
	router.SetTrustedProxies([]string{allowedHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	v1 := router.Group("/v1")

	v1NonAuthenticatedRouter := v1.Group("/")
	v1NonAuthenticatedRouter.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	authenticate := adapter.Wrap(auth0.EnsureValidToken().CheckJWT)
	v1AuthenticatedRouter := v1.Group("/", authenticate)
	{
		// this is a health check endpoint which is protected by the auth0 middleware
		v1AuthenticatedRouter.GET("/auth/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })
		UsersRoutes(v1AuthenticatedRouter, repository.NewUserRepo(gorm))
	}

	return router
}
