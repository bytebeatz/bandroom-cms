package dto

import (
	"time"

	"github.com/bytebeatz/bandroom-cms/core/model"
)

// CourseRequest defines the JSON body for creating/updating courses.
type CourseRequest struct {
	Slug        string         `json:"slug"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Language    string         `json:"language"`
	Difficulty  int            `json:"difficulty"`
	IsPublished bool           `json:"is_published"`
	Tags        []string       `json:"tags"`
	Metadata    map[string]any `json:"metadata"`
}

// CourseResponse defines the JSON returned by course endpoints.
type CourseResponse struct {
	ID          string         `json:"id"`
	Slug        string         `json:"slug"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Language    string         `json:"language"`
	Difficulty  int            `json:"difficulty"`
	IsPublished bool           `json:"is_published"`
	Tags        []string       `json:"tags,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	CreatorID   *string        `json:"creator_id,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// ToModel converts CourseRequest to model.Course.
func (r CourseRequest) ToModel() model.Course {
	return model.Course{
		Slug:        r.Slug,
		Title:       r.Title,
		Description: r.Description,
		Language:    r.Language,
		Difficulty:  model.DifficultyLevel(r.Difficulty),
		IsPublished: r.IsPublished,
		Tags:        r.Tags,
		Metadata:    r.Metadata,
	}
}

// FromModel maps model.Course to CourseResponse.
func FromModel(c model.Course) CourseResponse {
	var creatorID *string
	if c.CreatorID != nil {
		id := c.CreatorID.String()
		creatorID = &id
	}

	return CourseResponse{
		ID:          c.ID.String(),
		Slug:        c.Slug,
		Title:       c.Title,
		Description: c.Description,
		Language:    c.Language,
		Difficulty:  int(c.Difficulty),
		IsPublished: c.IsPublished,
		Tags:        c.Tags,
		Metadata:    c.Metadata,
		CreatorID:   creatorID,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

