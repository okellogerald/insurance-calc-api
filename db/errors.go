package db

import (
	"errors"

	"gorm.io/gorm"
)

const (
	ERR_RECORD_NOT_FOUND = "record not found"
)

// Looks for pre-defined errors and assigns them our preferred messages.
// The preferred errors are defined in /database/errors.go
func HandleDBError(r *gorm.DB) error {
	if r.Error != nil {
		if r.Error.Error() == "record not found" {
			return errors.New(ERR_RECORD_NOT_FOUND)
		}
		return r.Error
	}

	return nil
}
