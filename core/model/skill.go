package model

import (
	"time"

	"github.com/google/uuid"
)

type Skill struct {
	ID                   uuid.UUID      `json:"id"`
	CourseID             uuid.UUID      `json:"course_id"`
	UnitID               uuid.UUID      `json:"unit_id"`
	Slug                 string         `json:"slug"`
	Title                string         `json:"title"`
	Icon                 string         `json:"icon"`
	OrderIndex           int            `json:"order_index"`
	Difficulty           int            `json:"difficulty"`
	MaxCrowns            int            `json:"max_crowns"`
	BaseXPReward         int            `json:"base_xp_reward"`
	XPPerCrown           int            `json:"xp_per_crown"`
	PrerequisiteSkillIDs []uuid.UUID    `json:"prerequisite_skill_ids"`
	CreatorID            uuid.UUID      `json:"creator_id"`
	Tags                 []string       `json:"tags"`
	Metadata             map[string]any `json:"metadata"`
	Version              int            `json:"version"`
	DeletedAt            *time.Time     `json:"deleted_at,omitempty"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
}

