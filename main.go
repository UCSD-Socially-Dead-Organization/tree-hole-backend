package main

import (
	"time"

	env "github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/config"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/database"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/logger"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/repository"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/routers"
	"github.com/spf13/viper"
)

func main() {
	//set timezone. Default is Asia/Taipei Might need to change this to UTC
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Taipei")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	var err error
	var conf *env.Configuration

	// SetupConfig() will read the .env file and setup ALL the environment variables
	if conf, err = env.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	var gorm *database.GormDatabase
	if gorm, err = database.DBConnection(conf.DB.GetDSN(), conf); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}
	repository.InitDB(gorm)

	router := routers.Register(gorm, conf)
	logger.Fatalf("%v", router.Run(conf.Server.GenerateServerAddress()))
}
