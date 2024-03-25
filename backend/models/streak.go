package models

import (
	"time"

	"github.com/google/uuid"
)

type Streak struct {
	ID                  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key;unique"`
	StreakStart         time.Time `gorm:"autoCreateTime"`
	StreakEnd           *time.Time
	StreakLength        int        `gorm:"not null;default:1"`
	IsUserCurrentStreak bool       `gorm:"not null;default:true"`
	LastStreakDate      time.Time  `gorm:"not null;default:now()"`
	PreviousStreakID    *uuid.UUID `gorm:"type:uuid;"`
	UserID              string
}
