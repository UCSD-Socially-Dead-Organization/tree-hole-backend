package database

import (
	"log"

	env "github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDatabase struct {
	DB *gorm.DB
}

func DBConnection(masterDSN string, conf *env.Configuration) (*GormDatabase, error) {
	var db *gorm.DB
	var err error
	is_logMode := conf.DB.LogMode
	// TODO: Might need to add a debug flag
	// is_debug := viper.GetBool("DEBUG")

	loglevel := logger.Silent
	if is_logMode {
		loglevel = logger.Info
	}

	db, err = gorm.Open(postgres.Open(masterDSN), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})
	// TODO: Might need to handle replica
	// if !debug {
	// 	db.Use(dbresolver.Register(dbresolver.Config{
	// 		Replicas: []gorm.Dialector{
	// 			postgres.Open(replicaDSN),
	// 		},
	// 		Policy: dbresolver.RandomPolicy{},
	// 	}))
	// }

	if err != nil {
		log.Println("Db connection error")
		return nil, err
	}
	return &GormDatabase{DB: db}, nil
}
