package domain

import "time"

type MoviesSchedule struct {
	MovieName   string
	MovieGenre  string
	TheatreName string
	ShowTime    time.Time
}
