package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	ProfilePic []byte    `json:"profilePic" gorm:"type:bytea"`
	Username   string    `json:"username"`
	LastLogin  time.Time `json:"lastLogin"`
	Age        int       `json:"age"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
