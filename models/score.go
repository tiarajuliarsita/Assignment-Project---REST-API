package models

import "time"

type Score struct {
	ID              int    `json:"id" gorm:"primaryKey"`
	AssignmentTitle string `json:"assignment_title" gorm:"not null"`
	Score           int    `json:"score" gorm:"not null"`
	Description     string
	StudentID       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
