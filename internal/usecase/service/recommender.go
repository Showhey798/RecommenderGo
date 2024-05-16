package service

import (
	"context"

	"recommender.package/internal/domain/repository"
)

type Popularity struct {
	ClickRepo repository.ClickRepository
}

func (s *Popularity) GetPopularMovies(ctx context.Context) ([]uint32, error) {
	// get popular movies
	return s.ClickRepo.GetPopularMovies(ctx)
}
