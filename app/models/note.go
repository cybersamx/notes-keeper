package models

import "time"

type Note struct {
	ID uint	`gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Text	  string
}

