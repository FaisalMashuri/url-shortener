package repository

import (
	"backend/internal/domain/url"
	"context"

	"gorm.io/gorm"
)

type urlRepository struct {
	db *gorm.DB
}

// UpdateViewShortUrl implements url.UrlRepository.
func (repo *urlRepository) UpdateViewShortUrl(ctx context.Context, data url.Url) error {
	err := repo.db.Debug().WithContext(ctx).Model(&data).UpdateColumn("view", data.View).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUrlByShortUrl implements url.UrlRepository.
func (repo *urlRepository) GetUrlByShortUrl(ctx context.Context, shortUrl string) (url.Url, error) {
	var data url.Url
	err := repo.db.Debug().WithContext(ctx).First(&data, "short_url = ? ", shortUrl).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetAllUrl implements url.UrlRepository.
func (u *urlRepository) GetAllUrl(ctx context.Context) ([]url.Url, error) {
	var dataUrl []url.Url
	err := u.db.Debug().WithContext(ctx).Order("created_at DESC").Find(&dataUrl).Error
	if err != nil {
		return dataUrl, err
	}
	return dataUrl, nil
}

// CreateShortUrl implements url.UrlRepository.
func (repo *urlRepository) CreateShortUrl(ctx context.Context, data *url.Url) error {
	err := repo.db.Debug().WithContext(ctx).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUrlRepository(db *gorm.DB) url.UrlRepository {
	return &urlRepository{
		db: db,
	}
}
