package models

import (
	"time"

	"github.com/google/uuid"
)

type Match struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	User1     string    `json:"user1"`
	User2     string    `json:"user2"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
