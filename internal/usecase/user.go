package usecase

import (
	"context"

	domainmodel "recommender.package/internal/domain/entity"
	domainrepository "recommender.package/internal/domain/repository"
)

type UserUsecase struct {
	UserRepo domainrepository.UserRepository
}

func NewUserUsecase(userRepo domainrepository.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepo: userRepo,
	}
}

func (u *UserUsecase) GetUserByID(ctx context.Context, UserId string) (domainmodel.User, error) {
	user, err := u.UserRepo.GetByID(ctx, UserId)
	if err != nil {
		return domainmodel.User{}, err
	}
	return user, nil
}

func (u *UserUsecase) CreateUser(ctx context.Context, user domainmodel.User) error {
	err := u.UserRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
