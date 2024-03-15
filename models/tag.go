package models

import "time"

type Tags struct {
	TagID     int       `json:"tagId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}