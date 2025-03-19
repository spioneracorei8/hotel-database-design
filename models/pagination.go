package models

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalRows  int `json:"total_rows"`
	TotalPages int `json:"total_pages"`
}

func (p *Pagination) getOffset() int {
	return (p.Page - 1) * p.PerPage
}

func Paginate(value interface{}, where string, valArgs []interface{}, paginator *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Where(where, valArgs...).Count(&totalRows)
	paginator.TotalRows = int(totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(paginator.PerPage)))
	paginator.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(paginator.getOffset()).Limit(paginator.PerPage)
	}
}
