package repository

import (
	"backend/internal/domain/url"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// CreateShortUrl implements url.UrlRepository.
func (repo *userRepository) CreateShortUrl(ctx context.Context, data *url.Url) error {
	err := repo.db.Debug().WithContext(ctx).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUrlRepository(db *gorm.DB) url.UrlRepository {
	return &userRepository{
		db: db,
	}
}
