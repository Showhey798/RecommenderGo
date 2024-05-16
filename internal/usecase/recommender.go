package usecase

import (
	"context"

	"recommender.package/internal/domain/entity"
	"recommender.package/internal/domain/repository"
	"recommender.package/internal/usecase/service"
)

type RecommenderUsecase struct {
	MovieRepo repository.MovieRepository
	Services  service.Services
}

func NewRecommenderUsecase(movieRepo repository.MovieRepository, services service.Services) *RecommenderUsecase {
	return &RecommenderUsecase{
		MovieRepo: movieRepo,
		Services:  services,
	}
}

func (u *RecommenderUsecase) GetPopularModules(ctx context.Context, userId string) ([]entity.Movie, error) {
	movieIds, err := u.Services.Popularity.GetPopularMovies(ctx)
	if err != nil {
		return []entity.Movie{}, err
	}
	movie, err := u.MovieRepo.GetByIds(ctx, movieIds)
	if err != nil {
		return []entity.Movie{}, err
	}
	return movie, nil
}
