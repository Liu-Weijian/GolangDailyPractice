package repository

import (
	"GoProject/webook/internal/domain"
	"GoProject/webook/internal/repository/dao"
	"context"
)

type UserRepository struct {
	dao *dao.UserDao
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) (string, error) {

	return r.dao.Insert(ctx, domain.User{
		Email:    u.Email,
		Password: u.Password,
	})
}
