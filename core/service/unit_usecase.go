package service

import (
	"context"
	"fmt"
	"time"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/bytebeatz/bandroom-cms/core/repository"
	"github.com/google/uuid"
)

// UnitService handles business logic for units.
type UnitService struct {
	repo repository.UnitRepository
}

// NewUnitService initializes a new UnitService.
func NewUnitService(repo repository.UnitRepository) *UnitService {
	return &UnitService{repo: repo}
}

// CreateUnit handles creation logic including UUIDs, timestamps, versioning.
func (s *UnitService) CreateUnit(ctx context.Context, unit *model.Unit) error {
	unit.ID = uuid.New()
	unit.CreatedAt = time.Now().UTC()
	unit.UpdatedAt = unit.CreatedAt

	if unit.Version == 0 {
		unit.Version = 1
	}

	fmt.Printf("Creating unit: %+v\n", unit)
	return s.repo.Create(ctx, unit)
}

// UpdateUnit handles updating unit metadata and versioning.
func (s *UnitService) UpdateUnit(ctx context.Context, updated *model.Unit) error {
	existing, err := s.repo.GetByID(ctx, updated.ID)
	if err != nil {
		return fmt.Errorf("unit not found: %w", err)
	}

	updated.CreatedAt = existing.CreatedAt
	updated.UpdatedAt = time.Now().UTC()
	updated.Version = existing.Version + 1

	return s.repo.Update(ctx, updated)
}

// DeleteUnit deletes a unit by ID.
func (s *UnitService) DeleteUnit(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

// GetUnitByID fetches a unit by UUID.
func (s *UnitService) GetUnitByID(ctx context.Context, id uuid.UUID) (*model.Unit, error) {
	return s.repo.GetByID(ctx, id)
}

// ListUnitsByCourseID returns all units for a given course.
func (s *UnitService) ListUnitsByCourseID(
	ctx context.Context,
	courseID uuid.UUID,
) ([]*model.Unit, error) {
	return s.repo.ListByCourseID(ctx, courseID)
}
