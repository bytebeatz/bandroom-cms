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

// LessonService handles business logic for lessons.
type LessonService struct {
	repo repository.LessonRepository
}

func NewLessonService(repo repository.LessonRepository) *LessonService {
	return &LessonService{repo: repo}
}

func (s *LessonService) CreateLesson(
	ctx context.Context,
	lesson *model.Lesson,
	skillID uuid.UUID,
) error {
	exists, err := s.repo.ExistsByTitleInSkill(ctx, skillID, lesson.Title)
	if err != nil {
		return fmt.Errorf("error checking for duplicate lesson title: %w", err)
	}
	if exists {
		return fmt.Errorf("lesson with title '%s' already exists in this skill", lesson.Title)
	}

	now := time.Now().UTC()
	lesson.ID = uuid.New()
	lesson.SkillID = skillID
	lesson.CreatedAt = now
	lesson.UpdatedAt = now

	if strings.TrimSpace(lesson.Slug) == "" {
		lesson.Slug = utils.GenerateSlug(lesson.Title)
	}

	if lesson.Version == 0 {
		lesson.Version = 1
	}

	return s.repo.Create(ctx, lesson)
}

func (s *LessonService) UpdateLesson(ctx context.Context, updated *model.Lesson) error {
	existing, err := s.repo.GetByID(ctx, updated.ID)
	if err != nil {
		return fmt.Errorf("lesson not found: %w", err)
	}

	updated.CreatedAt = existing.CreatedAt
	updated.UpdatedAt = time.Now().UTC()
	updated.Version = existing.Version + 1

	if strings.TrimSpace(updated.Slug) == "" {
		updated.Slug = utils.GenerateSlug(updated.Title)
	}

	return s.repo.Update(ctx, updated)
}

func (s *LessonService) DeleteLesson(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *LessonService) GetLessonByID(ctx context.Context, id uuid.UUID) (*model.Lesson, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *LessonService) ListLessonsBySkillID(
	ctx context.Context,
	skillID uuid.UUID,
) ([]*model.Lesson, error) {
	return s.repo.ListBySkillID(ctx, skillID)
}

