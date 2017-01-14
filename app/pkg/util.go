package pkg

import "github.com/google/uuid"

// IsUUID ...
func IsUUID(id string) bool {
	if _, err := uuid.Parse(id); err == nil {
		return true
	}
	return false
}
