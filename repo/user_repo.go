package repo

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/udayangaac/shipments-service/repo/entity"
)

// UserRepo handles transactions related to the user.
type UserRepo interface {
	// FindByEmail function return user filter by email.
	FindByEmail(ctx context.Context, email string) (user entity.User, err error)
	// Save save user to the database.
	Save(ctx context.Context, user *entity.User) (err error)
}

// NewUserRepo create an instance of UserRepo implementation.
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		DB: db,
	}
}

type userRepo struct {
	DB *gorm.DB
}

// FindByEmail function return user filter by email.
func (u userRepo) FindByEmail(ctx context.Context, email string) (user entity.User, err error) {
	err = u.DB.Table(UsersTableName).Where("email = ?", email).First(&user).Error
	return
}

// Save save user to the database.
func (u userRepo) Save(ctx context.Context, user *entity.User) (err error) {
	err = u.DB.Table(UsersTableName).Create(user).Error
	return
}
