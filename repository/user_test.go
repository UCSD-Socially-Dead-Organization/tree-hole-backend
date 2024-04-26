package repository

import (
	"testing"
	"time"

	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/repository/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	gorm := *prepareDatabaseConnection(t)
	// drop all users
	gorm.DB.Exec("DELETE FROM users")

	userRepo := NewUserRepo(&gorm)

	uuid := uuid.New()
	now := time.Now()
	tests := []struct {
		name      string
		givenUser models.User
		wantUser  models.User
	}{
		{
			name: "Create a user",
			givenUser: models.User{
				ID:         uuid,
				ProfilePic: []byte("test"),
				Username:   "Jeffrey",
				LastLogin:  now,
				Age:        99,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			wantUser: models.User{
				ID:         uuid,
				ProfilePic: []byte("test"),
				Username:   "Jeffrey",
				LastLogin:  now,
				Age:        99,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
		}, // Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			err := userRepo.Create(&tt.givenUser)
			if err != nil {
				t.Errorf("Create() error = %v", err)
				return
			}
			var foundUser *models.User
			err = gorm.DB.First(&foundUser, tt.givenUser.ID).Error
			assert.NoError(t, err)
			assert.Equal(t, tt.wantUser.ID, foundUser.ID)
			assert.Equal(t, tt.wantUser.Age, foundUser.Age)
			assert.Equal(t, tt.wantUser.ProfilePic, foundUser.ProfilePic)
			assert.Equal(t, tt.wantUser.Username, foundUser.Username)
			assert.True(t, tt.wantUser.LastLogin.Equal(foundUser.LastLogin))
		})
	}
}

func Test_GetAll(t *testing.T) {
	gorm := *prepareDatabaseConnection(t)
	// drop all users
	gorm.DB.Exec("DELETE FROM users")

	userRepo := NewUserRepo(&gorm)

	now := time.Now()
	tests := []struct {
		name       string
		givenUsers []models.User
	}{
		{
			name: "Get all users",
			givenUsers: []models.User{
				{
					ID:         uuid.New(),
					ProfilePic: []byte("test"),
					Username:   "Jeffrey",
					LastLogin:  now,
					Age:        99,
					CreatedAt:  now,
					UpdatedAt:  now,
				},
				{
					ID:         uuid.New(),
					ProfilePic: []byte("test"),
					Username:   "Jeffrey",
					LastLogin:  now,
					Age:        99,
					CreatedAt:  now,
					UpdatedAt:  now,
				},
			},
		}, // Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, user := range tt.givenUsers {
				err := userRepo.Create(&user)
				if err != nil {
					t.Errorf("Create() error = %v", err)
					return
				}
			}
			users, err := userRepo.GetAll()
			assert.NoError(t, err)
			assert.Equal(t, len(tt.givenUsers), len(users))
		})
	}
}

func Test_GetOne(t *testing.T) {
	gorm := *prepareDatabaseConnection(t)
	// drop all users
	gorm.DB.Exec("DELETE FROM users")

	userRepo := NewUserRepo(&gorm)

	uuid := uuid.New()
	now := time.Now()
	tests := []struct {
		name      string
		givenUser models.User
	}{
		{
			name: "Get a user",
			givenUser: models.User{
				ID:         uuid,
				ProfilePic: []byte("test"),
				Username:   "Jeffrey",
				LastLogin:  now,
				Age:        99,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
		}, // Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userRepo.Create(&tt.givenUser)
			if err != nil {
				t.Errorf("Create() error = %v", err)
				return
			}
			user, err := userRepo.GetOne(tt.givenUser.ID)
			assert.NoError(t, err)
			assert.Equal(t, tt.givenUser.ID, user.ID)
			assert.Equal(t, tt.givenUser.Age, user.Age)
			assert.Equal(t, tt.givenUser.ProfilePic, user.ProfilePic)
			assert.Equal(t, tt.givenUser.Username, user.Username)
			assert.True(t, tt.givenUser.LastLogin.Equal(user.LastLogin))
		})
	}
}

func Test_Update(t *testing.T) {
	gorm := *prepareDatabaseConnection(t)
	// drop all users
	gorm.DB.Exec("DELETE FROM users")

	// create a user
	userRepo := NewUserRepo(&gorm)
	givenID := uuid.New()
	new_user := models.User{
		ID:         givenID,
		ProfilePic: []byte("test"),
		Username:   "Jeffrey",
		LastLogin:  time.Now(),
		Age:        99,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	err := userRepo.Create(&new_user)
	assert.NoError(t, err)

	now := time.Now()
	tests := []struct {
		name            string
		givenUserFields models.User
		wantUserFields  models.User
	}{
		{
			name: "Update a user",
			givenUserFields: models.User{
				ID:         givenID,
				ProfilePic: []byte("test"),
				Username:   "Jinhua",
				LastLogin:  now,
				Age:        30,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
		}, // Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userRepo.Update(&tt.givenUserFields)
			assert.NoError(t, err)
			user, err := userRepo.GetOne(tt.givenUserFields.ID)
			assert.NoError(t, err)
			assert.Equal(t, tt.givenUserFields.ID, user.ID)
			assert.Equal(t, tt.givenUserFields.ProfilePic, user.ProfilePic)
			assert.Equal(t, tt.givenUserFields.Username, user.Username)
			assert.True(t, tt.givenUserFields.LastLogin.Equal(user.LastLogin))
			assert.Equal(t, tt.givenUserFields.Age, user.Age)
		})
	}

}
