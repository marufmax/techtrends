package models

import (
	"context"
	"database/sql"
	"github.com/marufmax/techtrends/db"
	"github.com/uptrace/bun"
	"time"
)

type Category struct {
	bun.BaseModel `bun:"table:categories"`

	ID        uint64 `bun:",pk,autoincrement"`
	Name      string
	Type      string
	CreatedAt string
}

type JobCount struct {
	bun.BaseModel `bun:"table:job_counts"`

	ID           uint64 `bun:",pk,autoincrement"`
	Count        uint64
	Vacancy      uint64
	CategoryID   uint64
	Category     *Category `bun:"rel:belongs-to,join:category_id=id"`
	CategoryName string
	CategoryType string
	CrawlDate    time.Time `bun:",nullzero,notnull,default:current_date"`
	CreatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func GetAllCategories(ctx context.Context) ([]Category, error) {
	categories := make([]Category, 0)

	err := db.DB().NewSelect().Model(&categories).Where("is_active = ?", true).Limit(100).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func InsertManyJobs(ctx context.Context, jobCounts []JobCount) (sql.Result, error) {
	return db.DB().NewInsert().Model(&jobCounts).On("CONFLICT (category_id, crawl_date) DO UPDATE").Set("count=excluded.count, vacancy=excluded.vacancy").Exec(ctx)
}
