package repository

import (
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/database"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/repository/models"
)

/**
 * The repository package is responsible for handling the database operations. (Similar to the controllers in MVC but I call it repository)
 * It contains the database connection and the necessary functions to interact with the database and the models (tables).
 * By doing this, each model can have its own repository, which will handle the database operations for that model, and the handlers will only interact with the repositories.
 * This separation of concerns makes the code more modular and easier to maintain.
 * For example, if I want to test the database operations for the User model, I can create a mock repository that implements the UserRepository interface and use it in my tests.
 * This way I can test the "handlers" without actually interacting with the database.
 */

func InitDB(gorm *database.GormDatabase) error {
	// Create uuid function
	gorm.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	var err error
	db := gorm.DB

	// I dont think we need a separate file for migrations. We can just put the migrations in the SetUpDB function.
	err = db.AutoMigrate(&models.User{})
	err = db.AutoMigrate(&models.Match{})

	if err != nil {
		return err
	}

	return nil
}
