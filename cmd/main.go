package main

import (
	"fmt"
	"log"
	"net"

	"database/sql"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	controller "recommender.package/internal/adapter/grpc"
	"recommender.package/internal/infrastructure/postgres"
	pb "recommender.package/internal/proto/api"
	"recommender.package/internal/usecase"
	"recommender.package/internal/usecase/service"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db", 5432, "user", "password", "test")
	db, err := sql.Open(
		"postgres",
		dsn,
	)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	} else {
		log.Println("Database connected")
	}

	recommenderUsecase := usecase.NewRecommenderUsecase(
		&postgres.MovieRepository{DB: db},
		service.Services{
			Popularity: service.Popularity{
				ClickRepo: &postgres.ClickRepository{DB: db},
			},
		},
	)
	controller := controller.NewRecomenderController(recommenderUsecase)

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterRecommenderServiceServer(s, controller)

	log.Printf("Server started at :%v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
