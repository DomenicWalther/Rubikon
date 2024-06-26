package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Skin struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;unique"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Price         int            `json:"price" gorm:"not null;default:0"`
	Image         string         `json:"image" gorm:"not null;default:null"`
	Name          string         `json:"name" gorm:"not null;default:null"`
	Rarity        string         `json:"rarity" gorm:"not null;default:null"`
	RequiredLevel int            `json:"required_level" gorm:"not null;default:1"`
	Type          string         `json:"type" gorm:"not null;default:null"`
	Users         []User         `gorm:"many2many:user_skins;"`
}
