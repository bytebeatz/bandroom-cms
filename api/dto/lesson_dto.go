package dto

import (
	"time"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/bytebeatz/bandroom-cms/utils"
)

// LessonRequest defines the incoming JSON for creating or updating a lesson.
type LessonRequest struct {
	SkillID           string         `json:"skill_id"`
	Title             string         `json:"title"`
	Slug              string         `json:"slug"`
	Description       string         `json:"description,omitempty"`
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
	Tags              []string       `json:"tags"`
	Metadata          map[string]any `json:"metadata"`
}

// LessonResponse defines the JSON response for lesson data.
type LessonResponse struct {
	ID                string         `json:"id"`
	SkillID           string         `json:"skill_id"`
	Slug              string         `json:"slug"`
	Title             string         `json:"title"`
	Description       string         `json:"description,omitempty"`
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
	CreatorID         string         `json:"creator_id"`
	Tags              []string       `json:"tags"`
	Metadata          map[string]any `json:"metadata"`
	Version           int            `json:"version"`
	DeletedAt         *time.Time     `json:"deleted_at,omitempty"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

// ToModel converts a LessonRequest to model.Lesson.
func (r LessonRequest) ToModel() model.Lesson {
	return model.Lesson{
		SkillID:           utils.ParseUUID(r.SkillID),
		Title:             r.Title,
		Slug:              r.Slug,
		Description:       r.Description,
		OrderIndex:        r.OrderIndex,
		TotalExercises:    r.TotalExercises,
		BaseXP:            r.BaseXP,
		BonusXP:           r.BonusXP,
		RewardGems:        r.RewardGems,
		RewardHearts:      r.RewardHearts,
		RewardCondition:   r.RewardCondition,
		EstimatedDuration: r.EstimatedDuration,
		DifficultyRating:  r.DifficultyRating,
		IsTestable:        r.IsTestable,
		Tags:              r.Tags,
		Metadata:          r.Metadata,
	}
}

// FromLessonModel maps model.Lesson to LessonResponse.
func FromLessonModel(l model.Lesson) LessonResponse {
	return LessonResponse{
		ID:                l.ID.String(),
		SkillID:           l.SkillID.String(),
		Slug:              l.Slug,
		Title:             l.Title,
		Description:       l.Description,
		OrderIndex:        l.OrderIndex,
		TotalExercises:    l.TotalExercises,
		BaseXP:            l.BaseXP,
		BonusXP:           l.BonusXP,
		RewardGems:        l.RewardGems,
		RewardHearts:      l.RewardHearts,
		RewardCondition:   l.RewardCondition,
		EstimatedDuration: l.EstimatedDuration,
		DifficultyRating:  l.DifficultyRating,
		IsTestable:        l.IsTestable,
		CreatorID:         l.CreatorID.String(),
		Tags:              l.Tags,
		Metadata:          l.Metadata,
		Version:           l.Version,
		DeletedAt:         l.DeletedAt,
		CreatedAt:         l.CreatedAt,
		UpdatedAt:         l.UpdatedAt,
	}
}

