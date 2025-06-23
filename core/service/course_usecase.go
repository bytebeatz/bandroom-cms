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

// NewCourseService initializes a new CourseService.
func NewCourseService(repo repository.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

// CreateCourse handles creation logic including UUIDs, timestamps, slug generation, and context-driven user metadata.
func (s *CourseService) CreateCourse(ctx context.Context, course *model.Course) error {
	course.ID = uuid.New()
	course.CreatedAt = time.Now().UTC()
	course.UpdatedAt = course.CreatedAt

	// Generate slug from title if empty
	if strings.TrimSpace(course.Slug) == "" {
		course.Slug = utils.GenerateSlug(course.Title)
	}

	// Default version
	if course.Version == 0 {
		course.Version = 1
	}

	// Set creator ID from context (if present)
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

func (s *CourseService) UpdateCourse(ctx context.Context, course *model.Course) error {
	course.UpdatedAt = time.Now().UTC()
	course.Version += 1
	return s.repo.Update(ctx, course)
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

