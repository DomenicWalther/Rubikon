package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	OwnerID     string         `json:"owner_id" gorm:"not null"`
	ImageURL    string         `json:"imageURL"`
	UserCount   int            `json:"userCount" gorm:"default:1"`
	Users       []User         `gorm:"many2many:group_users;"`
	Is_Public   bool           `json:"isPrivate" gorm:"not null;default:true"`
	GroupChat   GroupChat      `gorm:"foreignkey:GroupID"`
}

type GroupChat struct {
	ID                uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();unique"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	GroupChatMessages []GroupChatMessage
	GroupID           uuid.UUID
}

type GroupChatMessage struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	GroupChatID uuid.UUID
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Message     string         `json:"message" gorm:"not null"`
	UserID      uuid.UUID      `json:"user_id" gorm:"not null"`
}
