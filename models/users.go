package models

import (
	"time"
)

type User struct {
	Id         int        `json:"id" gorm:"primary_key"`
	Username   string     `json:"username" binding:"required"`
	Age        int        `json:"age"`
	ProfilePic string     `json:"profile_pic"`
	LastLogin  *time.Time `json:"last_login,string,omitempty"`
	CreatedAt  *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at_at,string,omitempty"`
}

func (e *User) TableName() string {
	return "users"
}
