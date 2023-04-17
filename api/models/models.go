package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New().String()
	return
}

type BaseStatus string

const (
	Todo  BaseStatus = "todo"
	Doing BaseStatus = "doing"
	Done  BaseStatus = "done"
)

type Base struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ToDo struct {
	Base
	Title  string     `gorm:"not null" json:"title"`
	Detail string     `gorm:"not null" json:"detail"`
	Status BaseStatus `gorm:"not null" json:"status"`
}
