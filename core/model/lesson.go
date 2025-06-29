package model

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID                uuid.UUID      `json:"id"`
	SkillID           uuid.UUID      `json:"skill_id"`
	Slug              string         `json:"slug"`
	Title             string         `json:"title"`
	Description       string         `json:"description"`
	OrderIndex        int            `json:"order_index"`
	TotalExercises    int            `json:"total_exercises"`
	BaseXP            int            `json:"base_xp"`
	BonusXP           int            `json:"bonus_xp"`
	RewardGems        int            `json:"reward_gems"`
	RewardHearts      int            `json:"reward_hearts"`
	RewardCondition   string         `json:"reward_condition"`
	EstimatedDuration int            `json:"estimated_duration"`
	DifficultyRating  float32        `json:"difficulty_rating"`
	IsTestable        bool           `json:"is_testable"`
	CreatorID         uuid.UUID      `json:"creator_id"`
	Tags              []string       `json:"tags"`
	Metadata          map[string]any `json:"metadata"`
	Version           int            `json:"version"`
	DeletedAt         *time.Time     `json:"deleted_at,omitempty"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

