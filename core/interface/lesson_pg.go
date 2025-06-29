package _interface

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/bytebeatz/bandroom-cms/core/repository"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type lessonPG struct {
	db *sql.DB
}

func NewLessonPG(db *sql.DB) repository.LessonRepository {
	return &lessonPG{db: db}
}

func (r *lessonPG) Create(ctx context.Context, l *model.Lesson) error {
	metadataJSON, err := json.Marshal(l.Metadata)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO lessons (
			id, skill_id, slug, title, order_index, total_exercises, base_xp,
			bonus_xp, reward_gems, reward_hearts, reward_condition,
			estimated_duration, difficulty_rating, is_testable,
			creator_id, tags, metadata, version,
			deleted_at, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7,
			$8, $9, $10, $11,
			$12, $13, $14,
			$15, $16, $17, $18,
			$19, $20, $21
		)
	`

	_, err = r.db.ExecContext(ctx, query,
		l.ID, l.SkillID, l.Slug, l.Title, l.OrderIndex, l.TotalExercises, l.BaseXP,
		l.BonusXP, l.RewardGems, l.RewardHearts, l.RewardCondition,
		l.EstimatedDuration, l.DifficultyRating, l.IsTestable,
		l.CreatorID, pq.StringArray(l.Tags), metadataJSON, l.Version,
		l.DeletedAt, l.CreatedAt, l.UpdatedAt,
	)
	return err
}

func (r *lessonPG) Update(ctx context.Context, l *model.Lesson) error {
	metadataJSON, err := json.Marshal(l.Metadata)
	if err != nil {
		return err
	}

	query := `
		UPDATE lessons SET
			slug = $2, title = $3, order_index = $4, total_exercises = $5,
			base_xp = $6, bonus_xp = $7, reward_gems = $8, reward_hearts = $9,
			reward_condition = $10, estimated_duration = $11, difficulty_rating = $12,
			is_testable = $13, tags = $14, metadata = $15, version = $16,
			deleted_at = $17, updated_at = $18
		WHERE id = $1
	`

	_, err = r.db.ExecContext(ctx, query,
		l.ID, l.Slug, l.Title, l.OrderIndex, l.TotalExercises,
		l.BaseXP, l.BonusXP, l.RewardGems, l.RewardHearts,
		l.RewardCondition, l.EstimatedDuration, l.DifficultyRating,
		l.IsTestable, pq.StringArray(l.Tags), metadataJSON, l.Version,
		l.DeletedAt, l.UpdatedAt,
	)
	return err
}

func (r *lessonPG) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM lessons WHERE id = $1`, id)
	return err
}

func (r *lessonPG) GetByID(ctx context.Context, id uuid.UUID) (*model.Lesson, error) {
	query := `
		SELECT id, skill_id, slug, title, order_index, total_exercises, base_xp,
		       bonus_xp, reward_gems, reward_hearts, reward_condition,
		       estimated_duration, difficulty_rating, is_testable,
		       creator_id, tags, metadata, version,
		       deleted_at, created_at, updated_at
		FROM lessons WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	return scanLesson(row)
}

func (r *lessonPG) ListBySkillID(ctx context.Context, skillID uuid.UUID) ([]*model.Lesson, error) {
	query := `
		SELECT id, skill_id, slug, title, order_index, total_exercises, base_xp,
		       bonus_xp, reward_gems, reward_hearts, reward_condition,
		       estimated_duration, difficulty_rating, is_testable,
		       creator_id, tags, metadata, version,
		       deleted_at, created_at, updated_at
		FROM lessons WHERE skill_id = $1 ORDER BY order_index
	`
	rows, err := r.db.QueryContext(ctx, query, skillID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []*model.Lesson
	for rows.Next() {
		l, err := scanLesson(rows)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, l)
	}
	return lessons, nil
}

func (r *lessonPG) ExistsByTitleInSkill(
	ctx context.Context,
	skillID uuid.UUID,
	title string,
) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM lessons WHERE skill_id = $1 AND LOWER(title) = LOWER($2))`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, skillID, title).Scan(&exists)
	return exists, err
}

func scanLesson(scanner interface {
	Scan(dest ...any) error
}) (*model.Lesson, error) {
	var l model.Lesson
	var metadataBytes []byte

	err := scanner.Scan(
		&l.ID, &l.SkillID, &l.Slug, &l.Title, &l.OrderIndex, &l.TotalExercises, &l.BaseXP,
		&l.BonusXP, &l.RewardGems, &l.RewardHearts, &l.RewardCondition,
		&l.EstimatedDuration, &l.DifficultyRating, &l.IsTestable,
		&l.CreatorID, pq.Array(&l.Tags), &metadataBytes, &l.Version,
		&l.DeletedAt, &l.CreatedAt, &l.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	if len(metadataBytes) > 0 {
		if err := json.Unmarshal(metadataBytes, &l.Metadata); err != nil {
			return nil, err
		}
	}

	return &l, nil
}

