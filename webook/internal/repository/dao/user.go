package dao

import (
	"GoProject/webook/internal/domain"
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDao struct {
}

func NewUserDao(db *gorm.DB) *UserDao {

}

// User 直接对应数据库表结构
type User struct {
	Id       int64
	Email    string
	Password string
	//创建时间
	CreatedAt int64
	//更新时间
	UpdatedAt int64
}

func (dao *UserDao) Insert(ctx context.Context, u domain.User) (int64, error) {
	//存毫秒数
	now := time.Now().UnixMilli()

	return
}
