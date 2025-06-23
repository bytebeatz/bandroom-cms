package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/bytebeatz/bandroom-cms/core/repository"
	"github.com/bytebeatz/bandroom-cms/utils"
	"github.com/google/uuid"
)

// CourseService handles business logic for courses.
type CourseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) CreateCourse(ctx context.Context, course *model.Course) error {
	course.ID = uuid.New()
	course.CreatedAt = time.Now().UTC()
	course.UpdatedAt = course.CreatedAt

	if strings.TrimSpace(course.Slug) == "" {
		course.Slug = utils.GenerateSlug(course.Title)
	}

	if course.Version == 0 {
		course.Version = 1
	}

	if val := ctx.Value("user_id"); val != nil {
		if userIDStr, ok := val.(string); ok {
			if parsed, err := uuid.Parse(userIDStr); err == nil {
				course.CreatorID = &parsed
			}
		}
	}

	fmt.Printf("Creating course: %+v\n", course)
	return s.repo.Create(ctx, course)
}

func (s *CourseService) UpdateCourse(ctx context.Context, updated *model.Course) error {
	existing, err := s.repo.GetByID(ctx, updated.ID)
	if err != nil {
		return fmt.Errorf("course not found: %w", err)
	}

	updated.CreatedAt = existing.CreatedAt
	updated.UpdatedAt = time.Now().UTC()
	updated.Version = existing.Version + 1

	if updated.Slug == "" {
		updated.Slug = utils.GenerateSlug(updated.Title)
	}

	return s.repo.Update(ctx, updated)
}

func (s *CourseService) DeleteCourse(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *CourseService) GetCourseByID(ctx context.Context, id uuid.UUID) (*model.Course, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CourseService) GetCourseBySlug(ctx context.Context, slug string) (*model.Course, error) {
	return s.repo.GetBySlug(ctx, slug)
}

func (s *CourseService) ListCourses(
	ctx context.Context,
	publishedOnly bool,
) ([]*model.Course, error) {
	return s.repo.List(ctx, publishedOnly)
}

