package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key;unique"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	User_ID    string    `json:"id" gorm:"primary_key;unique"`
	Level      int       `json:"level" gorm:"not null;default:1"`
	Experience int       `json:"experience" gorm:"not null;default:0"`
	Currency   int       `json:"currency" gorm:"not null;default:0"`
	Skins      []Skin    `gorm:"many2many:user_skins;"`
	Streaks    []Streak  `gorm:"foreignKey:UserID;references:ID"`
}
