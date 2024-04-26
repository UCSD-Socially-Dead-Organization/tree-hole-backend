package models

import (
	"time"

	"github.com/google/uuid"
)

type Match struct {
	ID        uuid.UUID `gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	User1     uuid.UUID
	User2     uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
