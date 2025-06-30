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

// SkillService handles business logic for skills.
type SkillService struct {
	repo repository.SkillRepository
}

func NewSkillService(repo repository.SkillRepository) *SkillService {
	return &SkillService{repo: repo}
}

func (s *SkillService) CreateSkill(
	ctx context.Context,
	skill *model.Skill,
	courseID uuid.UUID,
) error {
	exists, err := s.repo.ExistsByTitleInCourse(ctx, courseID, skill.Title)
	if err != nil {
		return fmt.Errorf("error checking for duplicate skill title: %w", err)
	}
	if exists {
		return fmt.Errorf("skill with title '%s' already exists in this course", skill.Title)
	}

	now := time.Now().UTC()
	skill.ID = uuid.New()
	skill.CourseID = courseID
	skill.CreatedAt = now
	skill.UpdatedAt = now

	if strings.TrimSpace(skill.Slug) == "" {
		skill.Slug = utils.GenerateSlug(skill.Title)
	}

	if skill.Version == 0 {
		skill.Version = 1
	}

	// ✅ Validate creator_id is not nil (i.e. zero UUID)
	if skill.CreatorID == uuid.Nil {
		return fmt.Errorf("creator_id must be set")
	}

	return s.repo.Create(ctx, skill)
}

func (s *SkillService) UpdateSkill(ctx context.Context, updated *model.Skill) error {
	existing, err := s.repo.GetByID(ctx, updated.ID)
	if err != nil {
		return fmt.Errorf("skill not found: %w", err)
	}

	updated.CreatedAt = existing.CreatedAt
	updated.UpdatedAt = time.Now().UTC()
	updated.Version = existing.Version + 1

	if strings.TrimSpace(updated.Slug) == "" {
		updated.Slug = utils.GenerateSlug(updated.Title)
	}

	return s.repo.Update(ctx, updated)
}

func (s *SkillService) DeleteSkill(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *SkillService) GetSkillByID(ctx context.Context, id uuid.UUID) (*model.Skill, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *SkillService) ListSkillsByUnitID(
	ctx context.Context,
	unitID uuid.UUID,
) ([]*model.Skill, error) {
	return s.repo.ListByUnitID(ctx, unitID)
}
