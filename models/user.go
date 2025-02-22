package models

import "time"

type User struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	EmailVerified *time.Time `json:"emailVerified"`
	Image         string     `json:"image"`
}
