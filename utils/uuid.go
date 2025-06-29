package utils

import (
	"github.com/google/uuid"
)

// ParseUUID parses a single UUID string, returns zero UUID if invalid.
func ParseUUID(id string) uuid.UUID {
	parsed, _ := uuid.Parse(id)
	return parsed
}

// ParseUUIDs parses a slice of UUID strings into a slice of uuid.UUID.
func ParseUUIDs(ids []string) []uuid.UUID {
	var uuids []uuid.UUID
	for _, id := range ids {
		if parsed, err := uuid.Parse(id); err == nil {
			uuids = append(uuids, parsed)
		}
	}
	return uuids
}

// StringifyUUIDs converts a slice of uuid.UUID to a slice of strings.
func StringifyUUIDs(ids []uuid.UUID) []string {
	var strIDs []string
	for _, id := range ids {
		strIDs = append(strIDs, id.String())
	}
	return strIDs
}

