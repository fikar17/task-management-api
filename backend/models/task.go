package models

import (
	"time"

	"gorm.io/gorm"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in-progress"
	StatusCompleted  TaskStatus = "completed"
)

type Task struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"type:varchar(255);not null" json:"title" binding:"required,min=3"`
	Description string     `gorm:"type:text" json:"description"`
	DueDate     time.Time  `gorm:"type:date;not null" json:"due_date" binding:"required"`
	Status      TaskStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type CreateTaskInput struct {
	Title       string     `json:"title" binding:"required,min=3"`
	Description string     `json:"description"`
	DueDate     string     `json:"due_date" binding:"required"`
	Status      TaskStatus `json:"status"`
}

type UpdateTaskInput struct {
	Title       string     `json:"title" binding:"omitempty,min=3"`
	Description string     `json:"description"`
	DueDate     string     `json:"due_date"`
	Status      TaskStatus `json:"status"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Task{})
}
