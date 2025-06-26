package _interface

import (
	"context"
	"database/sql"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/bytebeatz/bandroom-cms/core/repository"
	"github.com/google/uuid"
)

type unitPG struct {
	db *sql.DB
}

// NewUnitPG returns a PostgreSQL-backed UnitRepository.
func NewUnitPG(db *sql.DB) repository.UnitRepository {
	return &unitPG{db: db}
}

func (r *unitPG) Create(ctx context.Context, u *model.Unit) error {
	query := `
		INSERT INTO units (
			id, course_id, title, description, "order_index", version,
			deleted_at, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9
		)
	`
	_, err := r.db.ExecContext(ctx, query,
		u.ID, u.CourseID, u.Title, u.Description, u.OrderIndex, u.Version,
		u.DeletedAt, u.CreatedAt, u.UpdatedAt,
	)
	return err
}

func (r *unitPG) Update(ctx context.Context, u *model.Unit) error {
	query := `
		UPDATE units SET
			title = $2,
			description = $3,
			"order_index" = $4,
			version = $5,
			deleted_at = $6,
			updated_at = $7
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		u.ID, u.Title, u.Description, u.OrderIndex, u.Version,
		u.DeletedAt, u.UpdatedAt,
	)
	return err
}

func (r *unitPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM units WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *unitPG) GetByID(ctx context.Context, id uuid.UUID) (*model.Unit, error) {
	query := `
		SELECT id, course_id, title, description, "order_index", version,
		       deleted_at, created_at, updated_at
		FROM units
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	return scanUnit(row)
}

func (r *unitPG) ListByCourseID(ctx context.Context, courseID uuid.UUID) ([]*model.Unit, error) {
	query := `
		SELECT id, course_id, title, description, "order_index", version,
		       deleted_at, created_at, updated_at
		FROM units
		WHERE course_id = $1
		ORDER BY "order_index"
	`
	rows, err := r.db.QueryContext(ctx, query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var units []*model.Unit
	for rows.Next() {
		u, err := scanUnit(rows)
		if err != nil {
			return nil, err
		}
		units = append(units, u)
	}
	return units, nil
}

func scanUnit(scanner interface {
	Scan(dest ...any) error
}) (*model.Unit, error) {
	var u model.Unit
	err := scanner.Scan(
		&u.ID,
		&u.CourseID,
		&u.Title,
		&u.Description,
		&u.OrderIndex,
		&u.Version,
		&u.DeletedAt,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

