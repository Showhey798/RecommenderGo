package postgres

import (
	"context"
	"time"

	"database/sql"

	"recommender.package/internal/domain/entity"
)

type MovieRepository struct {
	DB *sql.DB
}

func (repo *MovieRepository) Get(ctx context.Context, id uint32) (movie entity.Movie, err error) {
	row := repo.DB.QueryRowContext(ctx, "SELECT * FROM movies WHERE id = $1", id)
	err = row.Scan(&movie.ID, &movie.Title)
	if err != nil {
		return entity.Movie{}, err
	}
	return movie, nil
}

func (repo *MovieRepository) Create(ctx context.Context, movie entity.Movie) error {
	query := "INSERT INTO movie (id, title, created_on) VALUES ($1, $2, $3)"
	_, err := repo.DB.ExecContext(
		ctx, query, movie.ID, movie.Title, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (repo *MovieRepository) GetByIds(ctx context.Context, ids []uint32) ([]entity.Movie, error) {
	query := "SELECT * FROM movies WHERE id = ANY($1)"
	rows, err := repo.DB.QueryContext(ctx, query, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []entity.Movie
	for rows.Next() {
		var movie entity.Movie
		err := rows.Scan(&movie.ID, &movie.Title)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
