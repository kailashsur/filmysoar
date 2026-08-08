package utils

import "math"

const MaxPage = 100000

type Pagination struct {
	Page        int   `json:"page"`
	Limit       int   `json:"limit"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"totalPages"`
	HasNextPage bool  `json:"hasNextPage"`
	HasPrevPage bool  `json:"hasPrevPage"`
}

// NewPagination creates a pagination object
func NewPagination(page, limit int, total int64) *Pagination {
	page = NormalizePage(page)
	if limit < 1 {
		limit = 1
	}
	if total < 0 {
		total = 0
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &Pagination{
		Page:        page,
		Limit:       limit,
		Total:       total,
		TotalPages:  totalPages,
		HasNextPage: page < totalPages,
		HasPrevPage: page > 1,
	}
}

// NormalizePagination ensures client-provided pagination values are safe for
// database queries. Values outside the accepted range use the closest safe
// value rather than being passed through to GORM.
func NormalizePagination(page, limit, defaultLimit, maxLimit int) (int, int) {
	page = NormalizePage(page)

	if defaultLimit < 1 {
		defaultLimit = 1
	}
	if maxLimit < defaultLimit {
		maxLimit = defaultLimit
	}
	if limit < 1 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}

	return page, limit
}

// NormalizePage returns a page number that cannot create a negative or
// impractically large database offset.
func NormalizePage(page int) int {
	if page < 1 {
		return 1
	}
	if page > MaxPage {
		return MaxPage
	}
	return page
}

// GetOffset calculates the database offset for pagination
func GetOffset(page, limit int) int {
	page = NormalizePage(page)
	if limit < 1 {
		return 0
	}
	return (page - 1) * limit
}
