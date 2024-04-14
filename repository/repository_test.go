package repository

import (
	"testing"

	config "github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/config"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/database"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func prepareDatabaseConnection(t *testing.T) *database.GormDatabase {
	t.Helper()
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	assert.NoError(t, err)

	var conf *config.Configuration
	err = viper.Unmarshal(&conf)
	assert.NoError(t, err)

	var gorm *database.GormDatabase
	gorm, err = database.DBConnection(conf.DB.GetDSN(), conf)
	assert.NoError(t, err)

	InitDB(gorm)
	return gorm
}

func TestInitDB(t *testing.T) {
	prepareDatabaseConnection(t)

	// TODO: Add more tests here
	// For example, you can test if the tables are created successfully
}
