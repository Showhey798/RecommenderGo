package postgres

import (
	"context"
	"database/sql"
)

type ClickRepository struct {
	DB *sql.DB
}

func (repo *ClickRepository) GetPopularMovies(ctx context.Context) ([]uint32, error) {
	// top 10の人気アイテムを取得
	query := `
		SELECT movie_id
		FROM clicks
		GROUP BY movie_id
		ORDER BY COUNT(*) DESC
		LIMIT 10
	`
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var popularMovies []uint32
	for rows.Next() {
		var movieID uint32
		err := rows.Scan(&movieID)
		if err != nil {
			return nil, err
		}
		popularMovies = append(popularMovies, movieID)
	}
	return popularMovies, nil
}
