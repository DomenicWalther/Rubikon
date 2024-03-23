package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Streak struct {
	gorm.Model
	StreakStart         time.Time
	StreakEnd           *time.Time
	StreakLength        int  `gorm:"not null;default:0"`
	IsUserCurrentStreak bool `gorm:"not null;default:true"`
	LastStreakDate      time.Time
	UserID              uuid.UUID `gorm:"type:uuid;not null"`
}
