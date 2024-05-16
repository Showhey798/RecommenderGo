package repository

import "context"

type ClickRepository interface {
	GetPopularMovies(ctx context.Context) ([]uint32, error)
}
