package models

import "time"

type Review struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	MovieID   int       `json:"movie_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
