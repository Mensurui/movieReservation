package repository

import (
	"context"

	"github.com/Mensurui/movieReservation/internal/domain"
)

type AdminRepository interface {
	AddMovie(ctx context.Context, movie domain.Movie) error
	UpdateMovie(ctx context.Context, movie domain.Movie) error
	DeleteMovie(ctx context.Context, id int) error
	GetMovie(ctx context.Context) ([]domain.Movie, error)

	AddTheater(ctx context.Context, theater domain.Theater) error
	GetTheaterCapacity(ctx context.Context, name string) (int, error)

	AddMoviePremier(ctx context.Context, moviepremier domain.MoviePremier) error
}
