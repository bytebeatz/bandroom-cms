package repository

import (
	"context"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/google/uuid"
)

// LessonRepository defines contract for accessing lesson data.
type LessonRepository interface {
	Create(ctx context.Context, lesson *model.Lesson) error
	Update(ctx context.Context, lesson *model.Lesson) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Lesson, error)
	ListBySkillID(ctx context.Context, skillID uuid.UUID) ([]*model.Lesson, error)

	// For conflict checking
	ExistsByTitleInSkill(ctx context.Context, skillID uuid.UUID, title string) (bool, error)
}

