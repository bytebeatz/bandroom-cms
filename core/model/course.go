package model

import (
	"time"

	"github.com/google/uuid"
)

// DifficultyLevel defines difficulty categories for a course.
type DifficultyLevel int

const (
	Beginner     DifficultyLevel = 1
	Intermediate DifficultyLevel = 2
	Advanced     DifficultyLevel = 3
)

// Course represents a learning course in the system.
type Course struct {
	ID          uuid.UUID       `json:"id"`
	Slug        string          `json:"slug,omitempty"`
	Title       string          `json:"title"`
	Description string          `json:"description,omitempty"`
	Language    string          `json:"language,omitempty"`
	Difficulty  DifficultyLevel `json:"difficulty"`
	IsPublished bool            `json:"is_published"`

	Tags     []string               `json:"tags,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	Version   int        `json:"version"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	CreatorID *uuid.UUID `json:"creator_id,omitempty"`
}
