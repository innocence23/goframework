package model

import time "time"

type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Content   string    `json:"content"`
	Status    int       `json:"status"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
