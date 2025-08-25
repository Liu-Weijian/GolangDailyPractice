package service

import (
	"GoProject/webook/internal/domain"
	"GoProject/webook/internal/repository"
	"context"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Signup(ctx context.Context, u domain.User) error {
	//密码加密

	return svc.repo.Create(ctx, u)
}
