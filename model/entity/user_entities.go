package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	// Lists     []List `json:"lists"`
}
