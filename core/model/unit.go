package model

import (
	"time"

	"github.com/google/uuid"
)

type Unit struct {
	ID          uuid.UUID  `json:"id"`
	CourseID    uuid.UUID  `json:"course_id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	OrderIndex  int        `json:"order_index"`
	Version     int        `json:"version"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

