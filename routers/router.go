package routers

import (
	"net/http"

	config "github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/config"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/database"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/repository"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/routers/middleware"
	"github.com/gin-gonic/gin"
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
	router.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	// Set up the routes
	v1 := router.Group("/v1")
	{
		UsersRoutes(v1, repository.NewUserRepo(gorm))
	}

	return router
}
