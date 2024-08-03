package main

import (
	"github.com/Showhey798/RecommenderGo/cmd/config"
	gateway "github.com/Showhey798/RecommenderGo/internal/gateway/http"
	"github.com/Showhey798/RecommenderGo/internal/repository"
	"github.com/Showhey798/RecommenderGo/internal/repository/postgres"
	"github.com/Showhey798/RecommenderGo/internal/usecase"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	cfg := config.Config{
		Dsn:      "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
		DbDriver: "postgres",
	}
	db, err := postgres.SetUpDB(cfg.DbDriver, cfg.Dsn)

	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	database := repository.Database{
		Auth: &postgres.AuthRepository{DB: db},
	}
	uc := usecase.New(logger, database)
	gw := gateway.NewGateway(logger, uc)
	r := chi.NewRouter()
	gw.RegisterGateway(r)
	logger.Info("Server starting...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("failed to start server", zap.Error(err))
		panic(err)
	}
}
