package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	ProfilePic []byte    `gorm:"type:bytea"`
	Username   string
	LastLogin  time.Time
	Age        int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
