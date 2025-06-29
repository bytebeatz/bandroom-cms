package _interface

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/bytebeatz/bandroom-cms/core/model"
	"github.com/bytebeatz/bandroom-cms/core/repository"
	"github.com/bytebeatz/bandroom-cms/utils"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type skillPG struct {
	db *sql.DB
}

func NewSkillPG(db *sql.DB) repository.SkillRepository {
	return &skillPG{db: db}
}

func (r *skillPG) Create(ctx context.Context, s *model.Skill) error {
	query := `
		INSERT INTO skills (
			id, course_id, unit_id, slug, title, icon, order_index, difficulty,
			max_crowns, base_xp_reward, xp_per_crown, prerequisite_skill_ids,
			creator_id, tags, metadata, version,
			deleted_at, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8,
			$9, $10, $11, $12,
			$13, $14, $15, $16,
			$17, $18, $19
		)
	`

	// Fix array + JSON issues
	prereqs := pq.Array(utils.StringifyUUIDs(s.PrerequisiteSkillIDs))
	tags := pq.Array(s.Tags)
	jsonMetadata, err := json.Marshal(s.Metadata)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query,
		s.ID, s.CourseID, s.UnitID, s.Slug, s.Title, s.Icon, s.OrderIndex, s.Difficulty,
		s.MaxCrowns, s.BaseXPReward, s.XPPerCrown, prereqs,
		s.CreatorID, tags, jsonMetadata, s.Version,
		s.DeletedAt, s.CreatedAt, s.UpdatedAt,
	)
	return err
}

func (r *skillPG) Update(ctx context.Context, s *model.Skill) error {
	query := `
		UPDATE skills SET
			slug = $2, title = $3, icon = $4, order_index = $5, difficulty = $6,
			max_crowns = $7, base_xp_reward = $8, xp_per_crown = $9,
			prerequisite_skill_ids = $10, tags = $11, metadata = $12,
			version = $13, deleted_at = $14, updated_at = $15
		WHERE id = $1
	`

	prereqIDs := pq.StringArray(utils.StringifyUUIDs(s.PrerequisiteSkillIDs))
	tags := pq.StringArray(s.Tags)

	metadataJSON, err := json.Marshal(s.Metadata)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query,
		s.ID, s.Slug, s.Title, s.Icon, s.OrderIndex, s.Difficulty,
		s.MaxCrowns, s.BaseXPReward, s.XPPerCrown,
		prereqIDs, tags, metadataJSON,
		s.Version, s.DeletedAt, s.UpdatedAt,
	)
	return err
}

func (r *skillPG) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM skills WHERE id = $1`, id)
	return err
}

func (r *skillPG) GetByID(ctx context.Context, id uuid.UUID) (*model.Skill, error) {
	query := `
		SELECT id, course_id, unit_id, slug, title, icon, order_index, difficulty,
		       max_crowns, base_xp_reward, xp_per_crown, prerequisite_skill_ids,
		       creator_id, tags, metadata, version,
		       deleted_at, created_at, updated_at
		FROM skills WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	return scanSkill(row)
}

func (r *skillPG) ListByUnitID(ctx context.Context, unitID uuid.UUID) ([]*model.Skill, error) {
	query := `
		SELECT id, course_id, unit_id, slug, title, icon, order_index, difficulty,
		       max_crowns, base_xp_reward, xp_per_crown, prerequisite_skill_ids,
		       creator_id, tags, metadata, version,
		       deleted_at, created_at, updated_at
		FROM skills WHERE unit_id = $1 ORDER BY order_index
	`
	rows, err := r.db.QueryContext(ctx, query, unitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []*model.Skill
	for rows.Next() {
		s, err := scanSkill(rows)
		if err != nil {
			return nil, err
		}
		skills = append(skills, s)
	}
	return skills, nil
}

func (r *skillPG) ExistsByTitleInCourse(
	ctx context.Context,
	courseID uuid.UUID,
	title string,
) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM skills s
			JOIN units u ON s.unit_id = u.id
			WHERE u.course_id = $1 AND LOWER(s.title) = LOWER($2)
		)
	`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, courseID, title).Scan(&exists)
	return exists, err
}

func scanSkill(scanner interface {
	Scan(dest ...any) error
}) (*model.Skill, error) {
	var s model.Skill
	var prereqIDs pq.StringArray
	var tags pq.StringArray
	var metadataBytes []byte

	err := scanner.Scan(
		&s.ID, &s.CourseID, &s.UnitID, &s.Slug, &s.Title, &s.Icon, &s.OrderIndex, &s.Difficulty,
		&s.MaxCrowns, &s.BaseXPReward, &s.XPPerCrown, &prereqIDs,
		&s.CreatorID, &tags, &metadataBytes, &s.Version,
		&s.DeletedAt, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	s.PrerequisiteSkillIDs = utils.ParseUUIDs(prereqIDs)
	s.Tags = tags
	_ = json.Unmarshal(metadataBytes, &s.Metadata)

	return &s, nil
}

