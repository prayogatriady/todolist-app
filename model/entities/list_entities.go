package model

import (
	"time"

	"gorm.io/gorm"
)

type List struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int64  `json:"user_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	// Tasks       []Task `json:"tasks"`
}
