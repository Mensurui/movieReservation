package service

import (
	"context"

	"github.com/Mensurui/movieReservation/internal/domain"
	adminRepository "github.com/Mensurui/movieReservation/internal/repository/admin"
)

type AdminService struct {
	adminRepo adminRepository.AdminRepository
}

func NewAdminService(adminRepo adminRepository.AdminRepository) *AdminService {
	return &AdminService{
		adminRepo: adminRepo,
	}
}

func (as *AdminService) AddMovie(ctx context.Context, movie domain.Movie) error {
	err := as.adminRepo.AddMovie(ctx, movie)
	if err != nil {
		return err
	}
	return nil
}

func (as *AdminService) UpdateMovie(ctx context.Context, movie domain.Movie) error {
	err := as.adminRepo.UpdateMovie(ctx, movie)
	if err != nil {
		return err
	}
	return nil
}

func (as *AdminService) DeleteMovie(ctx context.Context, id int) error {
	err := as.adminRepo.DeleteMovie(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (as *AdminService) GetMovie(ctx context.Context) ([]domain.Movie, error) {
	movie, err := as.adminRepo.GetMovie(ctx)
	if err != nil {
		return []domain.Movie{}, err
	}

	return movie, nil
}

func (as *AdminService) AddTheater(ctx context.Context, theater domain.Theater) error {
	err := as.adminRepo.AddTheater(ctx, theater)
	if err != nil {
		return err
	}

	return nil
}

func (as *AdminService) GetTheaterCapacity(ctx context.Context, name string) (int, error) {
	capacity, err := as.adminRepo.GetTheaterCapacity(ctx, name)
	if err != nil {
		return 0, err
	}
	return capacity, nil
}

func (as *AdminService) AddMoviePremier(ctx context.Context, moviePremier domain.MoviePremier) error {
	err := as.adminRepo.AddMoviePremier(ctx, moviePremier)
	if err != nil {
		return err
	}

	return nil
}
