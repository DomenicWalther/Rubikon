package models

import (
	"time"

	"github.com/google/uuid"
)

type Streak struct {
	ID                  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key;unique"`
	StreakStart         time.Time
	StreakEnd           *time.Time
	StreakLength        int  `gorm:"not null;default:0"`
	IsUserCurrentStreak bool `gorm:"not null;default:true"`
	LastStreakDate      time.Time
	PreviousStreakID    *uuid.UUID `gorm:"type:uuid;"`
	UserID              string
}
