package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	Level      int       `json:"level" gorm:"not null;default:1"`
	Experience int       `json:"experience" gorm:"not null;default:0"`
	Currency   int       `json:"currency" gorm:"not null;default:0"`
	Skins      []Skin    `gorm:"many2many:user_skins;"`
	Streaks    []Streak
}
