package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"size:255"` 
	Content string `json:"content" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at" gorm:"index"`
	DeleteAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}