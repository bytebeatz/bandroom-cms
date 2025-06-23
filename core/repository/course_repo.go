package repository

import (
	"context"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/google/uuid"
)

// CourseRepository defines contract for accessing course data.
type CourseRepository interface {
	Create(ctx context.Context, course *model.Course) error
	Update(ctx context.Context, course *model.Course) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Course, error)
	GetBySlug(ctx context.Context, slug string) (*model.Course, error)
	List(ctx context.Context, publishedOnly bool) ([]*model.Course, error)
}
