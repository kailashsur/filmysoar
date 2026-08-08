package utils

import "testing"

func TestNormalizePagination(t *testing.T) {
	tests := []struct {
		name                string
		page, limit         int
		wantPage, wantLimit int
	}{
		{"defaults invalid values", 0, 0, 1, 20},
		{"normal values", 2, 50, 2, 50},
		{"caps page and limit", MaxPage + 1, 1000, MaxPage, 100},
		{"negative values", -4, -1, 1, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			page, limit := NormalizePagination(tt.page, tt.limit, 20, 100)
			if page != tt.wantPage || limit != tt.wantLimit {
				t.Fatalf("NormalizePagination(%d, %d) = (%d, %d), want (%d, %d)", tt.page, tt.limit, page, limit, tt.wantPage, tt.wantLimit)
			}
		})
	}
}

func TestNewPaginationHandlesInvalidLimit(t *testing.T) {
	pagination := NewPagination(0, 0, 10)
	if pagination.Page != 1 || pagination.Limit != 1 || pagination.TotalPages != 10 {
		t.Fatalf("unexpected pagination: %+v", pagination)
	}
}
