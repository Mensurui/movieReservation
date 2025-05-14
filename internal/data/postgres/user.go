package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Mensurui/movieReservation/internal/domain"
	repository "github.com/Mensurui/movieReservation/internal/repository/user"
	"github.com/lib/pq"
)

type postgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) repository.UserRepository {
	return &postgresUserRepository{
		db: db,
	}
}

func (pur *postgresUserRepository) Register(ctx context.Context, user domain.User) error {
	query := `
	INSERT INTO users(username)
	VALUES ($1)
	`

	_, err := pur.db.ExecContext(ctx, query, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func (pur *postgresUserRepository) GetMovies(ctx context.Context, dateTime time.Time) ([]domain.MoviesSchedule, error) {
	query := `
	SELECT
		m.name AS movie_name,
		m.genre AS movie_genre,
		th.hallname AS theater_name,
		mp.showtime
	FROM moviepremier mp
	JOIN movies m ON mp.movie_id = m.id
	JOIN theatre th ON mp.theatre_id = th.id
	WHERE showtime = $1;
	`

	var movieSchedules []domain.MoviesSchedule
	rows, err := pur.db.QueryContext(ctx, query, dateTime)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var movieSchedule domain.MoviesSchedule
		err := rows.Scan(
			&movieSchedule.MovieName,
			&movieSchedule.MovieGenre,
			&movieSchedule.TheatreName,
			&movieSchedule.ShowTime,
		)
		if err != nil {
			return nil, err
		}
		movieSchedules = append(movieSchedules, movieSchedule)
	}

	return movieSchedules, nil
}

func (pur *postgresUserRepository) ReserveSeat(ctx context.Context, moviePremierID, userID int) (err error) {
	tx, err := pur.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback() // rollback only if not committed
		}
	}()

	// Lock the row
	query := `
	SELECT current_capacity
	FROM reservationcapacity
	WHERE movie_premier_id = $1
	FOR UPDATE
	`
	var currentCap int
	err = tx.QueryRowContext(ctx, query, moviePremierID).Scan(&currentCap)
	if err != nil {
		return fmt.Errorf("failed to get current capacity: %w", err)
	}

	if currentCap == 0 {
		return fmt.Errorf("no seats available")
	}

	// Insert reservation
	insertQuery := `
	INSERT INTO reservations(user_id, movie_premier_id)
	VALUES($1, $2)
	`
	_, err = tx.ExecContext(ctx, insertQuery, userID, moviePremierID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return fmt.Errorf("user already reserved a seat")
		}
		return fmt.Errorf("insert failed: %w", err)
	}

	// Update capacity
	updateQuery := `
	UPDATE reservationcapacity
	SET current_capacity = $1
	WHERE movie_premier_id = $2
	`
	_, err = tx.ExecContext(ctx, updateQuery, currentCap-1, moviePremierID)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit failed: %w", err)
	}

	committed = true
	return nil
}
