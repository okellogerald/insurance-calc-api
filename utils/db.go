package utils

import (
	db "auth/db"
)

func IsRecordNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == db.ERR_RECORD_NOT_FOUND
}
