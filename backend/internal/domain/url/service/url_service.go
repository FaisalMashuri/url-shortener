package service

import (
	"backend/internal/domain/url"
	"backend/internal/domain/url/dto/request"
	"backend/internal/domain/url/dto/response"
	"context"
)

type urlService struct {
	repo url.UrlRepository
}

// GetUrlByShortUrl implements url.UrlService.
func (service *urlService) GetUrlByShortUrl(ctx context.Context, shortUrl string) (response.UrlResponse, error) {
	var res response.UrlResponse
	data, err := service.repo.GetUrlByShortUrl(ctx, shortUrl)
	if err != nil {
		return res, err
	}
	data.AddViewUrl()
	err = service.repo.UpdateViewShortUrl(ctx, data)
	if err != nil {
		return res, err
	}
	res = data.ToResponse()
	return res, nil
}

// GetAllUrl implements url.UrlService.
func (service *urlService) GetAllUrl(ctx context.Context) ([]response.UrlResponse, error) {
	var response []response.UrlResponse
	data, err := service.repo.GetAllUrl(ctx)
	if err != nil {
		return response, err
	}
	for _, url := range data {
		response = append(response, url.ToResponse())
	}
	return response, nil
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
