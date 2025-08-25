package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

// User 直接对应数据库表结构
type User struct {
	Id       int64  `gorm:"primaryKey,AUTO_INCREMENT"`
	Email    string `gorm:"unique"`
	Password string
	//创建时间
	CreatedAt int64
	//更新时间
	UpdatedAt int64
}

func (dao *UserDao) Insert(ctx context.Context, u User) error {
	//存毫秒数
	now := time.Now().UnixMilli()
	u.CreatedAt = now
	u.UpdatedAt = now
	return dao.db.WithContext(ctx).Create(&u).Error
}
