package model

import time "time"

type Product struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Code      string    `json:"code"`
	Price     int       `json:"price"`
}
