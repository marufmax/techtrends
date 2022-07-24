package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	gorm.Model
	JobCounts []JobCounts
	ID        uint `gorm:"primaryKey"`
	Name      string
	Type      string
	IsActive  bool
	CreatedAt time.Time
}

type JobCounts struct {
	gorm.Model
	Category     Category
	ID           uint `gorm:"primaryKey"`
	CategoryID   uint
	Count        int
	Vacancy      int
	CrawlDate    time.Time
	CategoryName string
	CategoryType string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
