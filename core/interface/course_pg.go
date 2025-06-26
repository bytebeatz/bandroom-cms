package _interface

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/bytebeatz/bandroom-cms/core/repository"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type coursePG struct {
	db *sql.DB
}

// NewCoursePG returns a PostgreSQL-backed CourseRepository.
func NewCoursePG(db *sql.DB) repository.CourseRepository {
	return &coursePG{db: db}
}

func (r *coursePG) Create(ctx context.Context, c *model.Course) error {
	tags := "{" + strings.Join(c.Tags, ",") + "}"
	meta, _ := json.Marshal(c.Metadata)

	query := `
	INSERT INTO courses (
		id, slug, title, description, language, difficulty, is_published,
		tags, metadata, version, deleted_at, created_at, updated_at, creator_id
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7,
		$8, $9, $10, $11, $12, $13, $14
	)
	`
	_, err := r.db.ExecContext(ctx, query,
		c.ID, c.Slug, c.Title, c.Description, c.Language, c.Difficulty, c.IsPublished,
		tags, meta, c.Version, c.DeletedAt, c.CreatedAt, c.UpdatedAt, c.CreatorID,
	)
	return err
}

func (r *coursePG) Update(ctx context.Context, c *model.Course) error {
	tags := "{" + strings.Join(c.Tags, ",") + "}"
	meta, _ := json.Marshal(c.Metadata)

	query := `
	UPDATE courses SET
		slug = $2, title = $3, description = $4, language = $5,
		difficulty = $6, is_published = $7, tags = $8, metadata = $9,
		version = $10, deleted_at = $11, updated_at = $12, creator_id = $13
	WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		c.ID, c.Slug, c.Title, c.Description, c.Language,
		c.Difficulty, c.IsPublished, tags, meta,
		c.Version, c.DeletedAt, c.UpdatedAt, c.CreatorID,
	)
	return err
}

func (r *coursePG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM courses WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *coursePG) GetByID(ctx context.Context, id uuid.UUID) (*model.Course, error) {
	query := `SELECT id, slug, title, description, language, difficulty, is_published, tags, metadata, version, deleted_at, created_at, updated_at, creator_id FROM courses WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	return scanCourse(row)
}

func (r *coursePG) GetBySlug(ctx context.Context, slug string) (*model.Course, error) {
	query := `SELECT id, slug, title, description, language, difficulty, is_published, tags, metadata, version, deleted_at, created_at, updated_at, creator_id FROM courses WHERE slug = $1`
	row := r.db.QueryRowContext(ctx, query, slug)
	return scanCourse(row)
}

func (r *coursePG) List(ctx context.Context, publishedOnly bool) ([]*model.Course, error) {
	query := `SELECT id, slug, title, description, language, difficulty, is_published, tags, metadata, version, deleted_at, created_at, updated_at, creator_id FROM courses`
	if publishedOnly {
		query += ` WHERE is_published = TRUE`
	}
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*model.Course
	for rows.Next() {
		c, err := scanCourse(rows)
		if err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

// ExistsByTitle checks if a course with the same title already exists (case-insensitive).
func (r *coursePG) ExistsByTitle(ctx context.Context, title string) (bool, error) {
	query := `SELECT EXISTS (
		SELECT 1 FROM courses WHERE LOWER(title) = LOWER($1)
	)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, title).Scan(&exists)
	return exists, err
}

// scanCourse extracts a Course from a DB row.
func scanCourse(scanner interface {
	Scan(dest ...any) error
}) (*model.Course, error) {
	var c model.Course
	var tags []string
	var metaRaw []byte

	err := scanner.Scan(
		&c.ID,
		&c.Slug,
		&c.Title,
		&c.Description,
		&c.Language,
		&c.Difficulty,
		&c.IsPublished,
		pq.Array(&tags),
		&metaRaw,
		&c.Version,
		&c.DeletedAt,
		&c.CreatedAt,
		&c.UpdatedAt,
		&c.CreatorID,
	)
	if err != nil {
		return nil, err
	}

	c.Tags = tags
	_ = json.Unmarshal(metaRaw, &c.Metadata)
	return &c, nil
}

