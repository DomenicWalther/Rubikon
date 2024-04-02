package models

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	ID          uint `json:"id" gorm:"primaryKey"`
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
}
