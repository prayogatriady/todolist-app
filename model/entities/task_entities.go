package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Completed   bool      `json:"completed"`
	ListID      int64     `json:"list_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
