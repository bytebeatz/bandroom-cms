package dto

import (
	"time"

	"github.com/bytebeatz/bandroom-cms/core/model"
)

type CourseRequest struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Difficulty  int    `json:"difficulty"`
	IsPublished bool   `json:"is_published"`
}

type CourseResponse struct {
	ID          string    `json:"id"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Difficulty  int       `json:"difficulty"`
	IsPublished bool      `json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (r CourseRequest) ToModel() model.Course {
	return model.Course{
		Slug:        r.Slug,
		Title:       r.Title,
		Description: r.Description,
		Language:    r.Language,
		Difficulty:  model.DifficultyLevel(r.Difficulty),
		IsPublished: r.IsPublished,
	}
}

func FromModel(c model.Course) CourseResponse {
	return CourseResponse{
		ID:          c.ID.String(),
		Slug:        c.Slug,
		Title:       c.Title,
		Description: c.Description,
		Language:    c.Language,
		Difficulty:  int(c.Difficulty),
		IsPublished: c.IsPublished,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}
