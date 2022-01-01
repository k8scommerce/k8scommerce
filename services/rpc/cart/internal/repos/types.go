package repos

import "math"

type PagingStats struct {
	TotalRecords int64 `db:"total_records" json:"total_records"`
	TotalPages   int64 `db:"total_pages" json:"total_pages"`
}

func (s *PagingStats) Calc(pageSize int64) *PagingStats {
	totalPages := float64(s.TotalRecords) / float64(pageSize)
	s.TotalPages = int64(math.Ceil(totalPages))
	return s
}
