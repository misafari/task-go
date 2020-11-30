package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name"`
	Surname   string         `json:"surname"`
	Username  string         `json:"username"`
	Tasks     []Task         `json:"tasks"`
}

type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Subject     string         `json:"subject"`
	Description string         `json:"description"`
	IsDone      bool           `json:"is_done"`
	UserID      uint           `json:"user_id"`
}
