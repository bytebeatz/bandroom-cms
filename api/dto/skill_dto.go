package dto

import (
	"time"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/bytebeatz/bandroom-cms/utils"
)

// SkillRequest defines the incoming JSON payload for creating or updating a skill.
type SkillRequest struct {
	CourseID             string         `json:"course_id"`              // Required
	UnitID               string         `json:"unit_id"`                // Required
	Title                string         `json:"title"`                  // Required
	Icon                 string         `json:"icon"`                   // Optional (default: 🎯)
	OrderIndex           int            `json:"order_index"`            // Required
	Difficulty           int            `json:"difficulty"`             // Required
	MaxCrowns            int            `json:"max_crowns"`             // Optional
	BaseXPReward         int            `json:"base_xp_reward"`         // Optional
	XPPerCrown           int            `json:"xp_per_crown"`           // Optional
	PrerequisiteSkillIDs []string       `json:"prerequisite_skill_ids"` // Optional
	Tags                 []string       `json:"tags"`                   // Optional
	Metadata             map[string]any `json:"metadata"`               // Optional
}

// SkillResponse defines the JSON response sent back to the client.
type SkillResponse struct {
	ID                   string         `json:"id"`
	CourseID             string         `json:"course_id"`
	UnitID               string         `json:"unit_id"`
	Title                string         `json:"title"`
	Slug                 string         `json:"slug"`
	Icon                 string         `json:"icon"`
	OrderIndex           int            `json:"order_index"`
	Difficulty           int            `json:"difficulty"`
	MaxCrowns            int            `json:"max_crowns"`
	BaseXPReward         int            `json:"base_xp_reward"`
	XPPerCrown           int            `json:"xp_per_crown"`
	PrerequisiteSkillIDs []string       `json:"prerequisite_skill_ids"`
	CreatorID            string         `json:"creator_id"`
	Tags                 []string       `json:"tags"`
	Metadata             map[string]any `json:"metadata"`
	Version              int            `json:"version"`
	DeletedAt            *time.Time     `json:"deleted_at,omitempty"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
}

// ToModel converts SkillRequest to model.Skill.
func (r SkillRequest) ToModel() model.Skill {
	prereqIDs := make([]string, len(r.PrerequisiteSkillIDs))
	copy(prereqIDs, r.PrerequisiteSkillIDs)

	return model.Skill{
		CourseID:             utils.ParseUUID(r.CourseID),
		UnitID:               utils.ParseUUID(r.UnitID),
		Title:                r.Title,
		Icon:                 r.Icon,
		OrderIndex:           r.OrderIndex,
		Difficulty:           r.Difficulty,
		MaxCrowns:            r.MaxCrowns,
		BaseXPReward:         r.BaseXPReward,
		XPPerCrown:           r.XPPerCrown,
		PrerequisiteSkillIDs: utils.ParseUUIDs(r.PrerequisiteSkillIDs),
		Tags:                 r.Tags,
		Metadata:             r.Metadata,
	}
}

// FromSkillModel maps a model.Skill to a SkillResponse.
func FromSkillModel(s model.Skill) SkillResponse {
	prereqs := utils.StringifyUUIDs(s.PrerequisiteSkillIDs)

	return SkillResponse{
		ID:                   s.ID.String(),
		CourseID:             s.CourseID.String(),
		UnitID:               s.UnitID.String(),
		Title:                s.Title,
		Slug:                 s.Slug,
		Icon:                 s.Icon,
		OrderIndex:           s.OrderIndex,
		Difficulty:           s.Difficulty,
		MaxCrowns:            s.MaxCrowns,
		BaseXPReward:         s.BaseXPReward,
		XPPerCrown:           s.XPPerCrown,
		PrerequisiteSkillIDs: prereqs,
		CreatorID:            s.CreatorID.String(),
		Tags:                 s.Tags,
		Metadata:             s.Metadata,
		Version:              s.Version,
		DeletedAt:            s.DeletedAt,
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
	}
}

