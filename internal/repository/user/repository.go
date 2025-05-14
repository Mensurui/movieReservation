package user

import (
	"context"
	"time"

	"github.com/Mensurui/movieReservation/internal/domain"
)

type UserRepository interface {
	Register(ctx context.Context, user domain.User) error
	GetMovies(ctx context.Context, dateTime time.Time) ([]domain.MoviesSchedule, error)
	ReserveSeat(ctx context.Context, movieID, userID int) error
}
