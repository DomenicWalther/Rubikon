package models

import (
	"time"

	"gorm.io/gorm"
)

type Streak struct {
	gorm.Model
	StreakStart         time.Time
	StreakEnd           *time.Time
	StreakLength        int  `gorm:"not null;default:0"`
	IsUserCurrentStreak bool `gorm:"not null;default:true"`
	LastStreakDate      time.Time
}
