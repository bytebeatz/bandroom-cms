package repository

import (
	"context"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/google/uuid"
)

type UnitRepository interface {
	Create(ctx context.Context, unit *model.Unit) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Unit, error)
	ListByCourseID(ctx context.Context, courseID uuid.UUID) ([]*model.Unit, error)
	Update(ctx context.Context, unit *model.Unit) error
	Delete(ctx context.Context, id uuid.UUID) error
}
