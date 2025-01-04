package models

import (
	"time"
)

type TodoItem struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id" form:"id"` 
	Title       string    `gorm:"size:255;not null" json:"title" form:"title" binding:"required"` 
	Description string    `gorm:"type:text" json:"description" form:"description"` 
	Status      string    `gorm:"size:50;not null;default:'pending'" json:"status" form:"status"` 
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"` 
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"` 
}

type TodoItemCreation struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
