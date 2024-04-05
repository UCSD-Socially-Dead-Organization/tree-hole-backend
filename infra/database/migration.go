package database

import (
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/models"
)

// Add list of model add for migrations
var migrationModels = []interface{}{&models.User{}}
