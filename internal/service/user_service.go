package service

import (
	"context"
	"log"
	"time"

	"github.com/Mensurui/movieReservation/internal/domain"
	repository "github.com/Mensurui/movieReservation/internal/repository/user"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) Register(ctx context.Context, user domain.User) error {
	err := us.userRepo.Register(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetMovie(ctx context.Context, dateTime time.Time) ([]domain.MoviesSchedule, error) {
	movies, err := us.userRepo.GetMovies(ctx, dateTime)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (us *UserService) ReserveSeat(ctx context.Context, moviePremierID, userID int) error {
	err := us.userRepo.ReserveSeat(ctx, moviePremierID, userID)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return nil
}
