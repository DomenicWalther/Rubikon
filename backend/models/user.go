package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	User_ID           string         `json:"id" gorm:"primaryKey;unique"`
	Level             int            `json:"level" gorm:"not null;default:1"`
	Experience        int            `json:"experience" gorm:"not null;default:0"`
	MonthlyExperience int            `json:"monthly_experience" gorm:"not null;default:0"`
	Currency          int            `json:"currency" gorm:"not null;default:0"`
	Skins             []Skin         `gorm:"many2many:user_skins;"`
	Streaks           []Streak       `gorm:"foreignKey:UserID;references:User_ID"`
	Groups            []Group        `gorm:"many2many:group_users;"`
}
