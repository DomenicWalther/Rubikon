package handlers

import (
	"testing"
	"time"

	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/stretchr/testify/assert"
)

func TestDetermineStreakAction(t *testing.T) {
	testCases := []struct {
		name     string
		streak   *models.Streak
		expected string
	}{
		{
			name: "Today",
			streak: &models.Streak{
				LastStreakDate: time.Now(),
			},
			expected: "Today",
		},
		{
			name: "Yesterday",
			streak: &models.Streak{
				LastStreakDate: time.Now().AddDate(0, 0, -1),
			},
			expected: "Yesterday",
		},
		{
			name: "Older",
			streak: &models.Streak{
				LastStreakDate: time.Now().AddDate(0, 0, -2),
			},
			expected: "Older",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := determineStreakAction(tc.streak)
			assert.Equal(t, tc.expected, result)
		})
	}
}
