package usecase

import (
	"context"

	"recommender.package/internal/domain/entity"
	"recommender.package/internal/domain/repository"
)

type MovieUsecase struct {
	MovieRepo repository.MovieRepository
}

func NewMovieUsecase(movieRepo repository.MovieRepository) *MovieUsecase {
	return &MovieUsecase{
		MovieRepo: movieRepo,
	}
}

func (u *MovieUsecase) GetUserByID(ctx context.Context, MovieId string) (entity.Movie, error) {
	movie, err := u.MovieRepo.Get(ctx, MovieId)
	if err != nil {
		return entity.Movie{}, err
	}
	return movie, nil
}

func (u *MovieUsecase) CreateMovie(ctx context.Context, movie entity.Movie) error {
	err := u.MovieRepo.Create(ctx, movie)
	if err != nil {
		return err
	}
	return nil
}

func (u *MovieUsecase) GetMoviesFromIds(ctx context.Context, ids []string) ([]entity.Movie, error) {
	movies, err := u.MovieRepo.GetByIds(ctx, ids)
	if err != nil {
		return []entity.Movie{}, err
	}
	return movies, nil
}
