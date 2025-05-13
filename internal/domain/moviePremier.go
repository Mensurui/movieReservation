package domain

import "time"

type MoviePremier struct {
	MovieID   int       `json:"movie_id"`
	ShowTime  time.Time `json:"show_time"`
	TheaterID int       `json:"theater_id"`
	Price     float64   `json:"price"`
}
