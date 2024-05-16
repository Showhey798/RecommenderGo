package controller

import (
	"context"

	grpcapi "recommender.package/internal/proto/api"
	"recommender.package/internal/usecase"
)

type RecommenderController struct {
	usecase *usecase.RecommenderUsecase
	grpcapi.RecommenderServiceServer
}

func NewRecomenderController(usecase *usecase.RecommenderUsecase) *RecommenderController {
	return &RecommenderController{
		usecase: usecase,
	}

}

func (handler *RecommenderController) GetPersonalizedModule(ctx context.Context, req *grpcapi.GetModulesRequest) (*grpcapi.GetModulesResponse, error) {

	movies, err := handler.usecase.GetPopularModules(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	moviesResponse := make([]*grpcapi.Movie, 0, len(movies))
	for _, movie := range movies {
		moviesResponse = append(moviesResponse, &grpcapi.Movie{Id: movie.ID, Title: movie.Title})
	}

	popularModule := &grpcapi.Module{
		Name:   "人気の映画",
		Movies: moviesResponse,
	}

	return &grpcapi.GetModulesResponse{
		Modules: []*grpcapi.Module{popularModule},
	}, nil
}
