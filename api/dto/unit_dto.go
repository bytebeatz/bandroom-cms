package dto

import (
	"time"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/google/uuid"
)

// UnitRequest defines the JSON body for creating/updating units.
type UnitRequest struct {
	CourseID    string `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	OrderIndex  int    `json:"order_index"`
}

// UnitResponse defines the JSON returned by unit endpoints.
type UnitResponse struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	OrderIndex  int       `json:"order_index"`
	Version     int       `json:"version"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ToModel converts UnitRequest to model.Unit.
func (r UnitRequest) ToModel() model.Unit {
	courseUUID, _ := uuid.Parse(r.CourseID)

	return model.Unit{
		CourseID:    courseUUID,
		Title:       r.Title,
		Description: r.Description,
		OrderIndex:  r.OrderIndex,
	}
}

// FromUnitModel maps model.Unit to UnitResponse.
func FromUnitModel(u model.Unit) UnitResponse {
	return UnitResponse{
		ID:          u.ID.String(),
		CourseID:    u.CourseID.String(),
		Title:       u.Title,
		Description: u.Description,
		OrderIndex:  u.OrderIndex,
		Version:     u.Version,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

