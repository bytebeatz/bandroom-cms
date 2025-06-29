package repository

import (
	"context"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/google/uuid"
)

// SkillRepository defines contract for accessing skill data.
type SkillRepository interface {
	Create(ctx context.Context, skill *model.Skill) error
	Update(ctx context.Context, skill *model.Skill) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Skill, error)
	ListByUnitID(ctx context.Context, unitID uuid.UUID) ([]*model.Skill, error)

	// For conflict checking
	ExistsByTitleInCourse(ctx context.Context, courseID uuid.UUID, title string) (bool, error)
}

