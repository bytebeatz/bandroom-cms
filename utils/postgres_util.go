package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// NullUUID helps with scanning nullable UUIDs from DB rows.
func NullUUID(u *uuid.UUID) any {
	if u == nil {
		return nil
	}
	return *u
}

// ScanUUID safely scans a UUID from a DB row.
func ScanUUID(src any) (*uuid.UUID, error) {
	switch v := src.(type) {
	case string:
		id, err := uuid.Parse(v)
		if err != nil {
			return nil, fmt.Errorf("invalid UUID string: %w", err)
		}
		return &id, nil
	case []byte:
		id, err := uuid.Parse(string(v))
		if err != nil {
			return nil, fmt.Errorf("invalid UUID bytes: %w", err)
		}
		return &id, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("unsupported UUID type: %T", src)
	}
}

// NullString helps insert optional strings into the database.
func NullString(s *string) sql.NullString {
	if s == nil || *s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

// MustJSON encodes input to JSON or panics (for internal use only).
func MustJSON(v any) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal JSON: %v", err))
	}
	return data
}
