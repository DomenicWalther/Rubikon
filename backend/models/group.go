package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID                uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();unique"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt     `gorm:"index"`
	Name              string             `json:"name" gorm:"not null"`
	Description       string             `json:"description"`
	OwnerID           string             `json:"owner_id" gorm:"not null"`
	ImageURL          string             `json:"imageURL"`
	UserCount         int                `json:"userCount" gorm:"default:1"`
	Users             []User             `gorm:"many2many:group_users;"`
	Is_Public         bool               `json:"isPrivate" gorm:"not null;default:true"`
	GroupChatMessages []GroupChatMessage `gorm:"foreignkey:GroupID"`
}

type GroupChatMessage struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	GroupID   uuid.UUID      `json:"group_id" gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Message   string         `json:"message" gorm:"not null"`
	UserID    string         `json:"user_id" gorm:"not null"`
}
