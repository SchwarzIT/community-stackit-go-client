package permissions

import "errors"

func validatePagination(limit, offset int) error {
	if offset < 0 {
		return errors.New("offset must be >= 0")
	}
	if limit < 1 {
		return errors.New("limit must be >= 1")
	}
	return nil
}
