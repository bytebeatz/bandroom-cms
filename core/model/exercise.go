package model

import (
	"time"

	"github.com/google/uuid"
)

// ExerciseType defines the type of exercise.
type ExerciseType string

const (
	ExerciseMultipleChoice   ExerciseType = "multiple_choice"
	ExerciseAudioRecognition ExerciseType = "audio_recognition"
	ExercisePlayback         ExerciseType = "playback"
	ExerciseFillInTheBlank   ExerciseType = "fill_in_the_blank"
	ExerciseMatching         ExerciseType = "matching"
	ExerciseTyping           ExerciseType = "typing"
)

// MatchingType defines the interaction format for matching exercises.
type MatchingType string

const (
	MatchTextToText   MatchingType = "text_to_text"
	MatchImageToText  MatchingType = "image_to_text"
	MatchAudioToText  MatchingType = "audio_to_text"
	MatchTextToAudio  MatchingType = "text_to_audio"
	MatchImageToImage MatchingType = "image_to_image"
	MatchAudioToAudio MatchingType = "audio_to_audio"
)

// Exercise represents a single learning task within a lesson or skill.
type Exercise struct {
	ID           uuid.UUID     `json:"id"`
	SkillID      uuid.UUID     `json:"skill_id"`
	LessonID     uuid.UUID     `json:"lesson_id"`
	Title        string        `json:"title"`
	Type         ExerciseType  `json:"type"`
	MatchingType *MatchingType `json:"matching_type,omitempty"` // only used for matching type
	Prompt       string        `json:"prompt"`                  // question or instruction shown to learner
	MediaURL     *string       `json:"media_url,omitempty"`     // optional media (image or audio)
	OrderIndex   int           `json:"order_index"`             // position in skill or lesson
	Points       int           `json:"points"`                  // XP or score value
	Grade        int           `json:"grade"`                   // music theory grade level (1–5)
	Syllabus     string        `json:"syllabus"`                // e.g. "ABRSM", "Trinity"
	ObjectiveTag string        `json:"objective"`               // e.g. "intervals", "notation", etc.
	Metadata     JSONB         `json:"metadata"`                // flexible config per exercise type
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at,omitempty"`
}

// ExerciseOption defines a selectable choice for MCQ or matching exercises.
type ExerciseOption struct {
	ID         uuid.UUID `json:"id"`
	ExerciseID uuid.UUID `json:"exercise_id"`
	Label      string    `json:"label"`       // display text
	Value      string    `json:"value"`       // internal answer value
	IsCorrect  bool      `json:"is_correct"`  // whether it's a correct option
	MediaURL   *string   `json:"media_url"`   // optional (e.g. for audio/image options)
	OrderIndex int       `json:"order_index"` // position in UI
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ExerciseAnswer stores the user's submitted response.
type ExerciseAnswer struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	ExerciseID  uuid.UUID `json:"exercise_id"`
	AttemptID   uuid.UUID `json:"attempt_id"` // ties to a LessonAttempt or SkillAttempt
	Response    JSONB     `json:"response"`   // user response (can be structured or freeform)
	IsCorrect   bool      `json:"is_correct"`
	SubmittedAt time.Time `json:"submitted_at"`
}

// JSONB is a flexible structure for storing arbitrary JSON.
type JSONB map[string]any
