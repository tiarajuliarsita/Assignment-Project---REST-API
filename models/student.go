package models

import "time"

type Student struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Age       int    `json:"age" gorm:"not null"`
	Scores    []Score
	CreatedAt time.Time
	UpdatedAt time.Time
}
