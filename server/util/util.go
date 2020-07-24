package util

import (
	"strings"

	"github.com/google/uuid"
)

// CreateUniqueID is
func CreateUniqueID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
