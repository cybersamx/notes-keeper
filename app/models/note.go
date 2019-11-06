package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type Note struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Text      string
}

// Override to use UUID than numeric ID.
func (note *Note) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4().String()

	return scope.SetColumn("ID", id)
}
