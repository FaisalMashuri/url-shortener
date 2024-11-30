package service

import (
	"backend/internal/domain/url"
	"backend/internal/domain/url/dto/request"
	"context"
)

type urlService struct {
	repo url.UrlRepository
}

// CreateShortUrl implements url.UrlService.
func (service *urlService) CreateShortUrl(ctx context.Context, data request.UrlRequest) (string, error) {
	dataUrl, err := url.NewUser(data)
	if err != nil {
		return "", err
	}
	errCreateUrl := service.repo.CreateShortUrl(ctx, &dataUrl)
	if errCreateUrl != nil {
		return "", errCreateUrl
	}
	return dataUrl.ShortUrl, nil

}

func NewUrlService(repo url.UrlRepository) url.UrlService {
	return &urlService{
		repo: repo,
	}
}
