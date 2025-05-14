package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/Mensurui/movieReservation/internal/domain"
	repository "github.com/Mensurui/movieReservation/internal/repository/admin"
)

type postgresAdminRepository struct {
	db *sql.DB
}

func NewPostgresAdminRepository(db *sql.DB) repository.AdminRepository {
	return &postgresAdminRepository{
		db: db,
	}
}

func (par *postgresAdminRepository) AddMovie(ctx context.Context, movie domain.Movie) error {
	log.Printf("Adding Movie")
	query := `
	INSERT INTO movies(name, genre)
	VALUES($1, $2)
	`
	_, err := par.db.ExecContext(ctx, query, movie.Name, movie.Genre)
	if err != nil {
		log.Printf("Error Adding Movie")
		return err
	}
	log.Printf("Finished Adding Movie")
	return nil
}

func (par *postgresAdminRepository) UpdateMovie(ctx context.Context, movie domain.Movie) error {
	if movie.Genre == "" {
		query := `
	UPDATE movies
	SET name = $1
	`
		_, err := par.db.ExecContext(ctx, query, movie.Name)
		if err != nil {
			return err
		}
		return nil
	}

	query := `
	UPDATE movies
	SEST genre = $1
	`
	_, err := par.db.ExecContext(ctx, query, movie.Genre)
	if err != nil {
		return err
	}

	return nil
}

func (par *postgresAdminRepository) DeleteMovie(ctx context.Context, id int) error {
	query := `
	DELETE
	FROM movies
	WHERE id = $1
	`

	_, err := par.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (par *postgresAdminRepository) GetMovie(ctx context.Context) ([]domain.Movie, error) {
	query := `
	SELECT name, genre
	FROM movies
	`

	rows, err := par.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []domain.Movie
	for rows.Next() {
		var movie domain.Movie
		err := rows.Scan(&movie.Name, &movie.Genre)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (par *postgresAdminRepository) AddTheater(ctx context.Context, theater domain.Theater) error {
	query := `
	INSERT INTO theatre(hallname, capacity)
	VALUES ($1, $2)
	`
	_, err := par.db.ExecContext(ctx, query, theater.HallName, theater.Capacity)
	if err != nil {
		return err
	}
	return nil
}

func (par *postgresAdminRepository) GetTheaterCapacity(ctx context.Context, name string) (int, error) {
	query := `
	SELECT capacity
	FROM theatre
	WHERE hallname = $1
	`
	var capacity int
	err := par.db.QueryRowContext(ctx, query, name).Scan(&capacity)
	if err != nil {
		return 0, err
	}
	return capacity, nil
}

func (par *postgresAdminRepository) AddMoviePremier(ctx context.Context, moviePremier domain.MoviePremier) error {
	tx, err := par.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()
	query := `
	INSERT INTO moviepremier(showtime, price, movie_id, theatre_id)
	VALUES ($1,$2,$3,$4)
	RETURNING id
	`

	var id int
	err = tx.QueryRowContext(ctx, query, moviePremier.ShowTime, moviePremier.Price, moviePremier.MovieID, moviePremier.TheaterID).Scan(&id)
	if err != nil {
		return err
	}

	var capacity int
	queryTheater := `
	SELECT capacity
	FROM theatre
	WHERE id = $1
	`
	err = tx.QueryRowContext(ctx, queryTheater, moviePremier.TheaterID).Scan(&capacity)
	if err != nil {
		return err
	}

	queryCapacity := `
	INSERT INTO reservationcapacity(movie_premier_id, current_capacity)
	VALUES($1, $2)
	`
	_, err = tx.ExecContext(ctx, queryCapacity, id, capacity)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	committed = true

	return nil
}
