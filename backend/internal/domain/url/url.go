package url

import (
	"backend/internal/domain/url/dto/request"
	"backend/internal/domain/url/dto/response"
	"context"

	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	OriginalUrl string
	ShortUrl    string
	View        int16
}

func NewUser(request request.UrlRequest) (Url, error) {
	shortURl, err := gonanoid.New()
	if err != nil {
		return Url{}, nil
	}
	return Url{
		OriginalUrl: request.Url,
		ShortUrl:    shortURl,
	}, nil
}

func (u *Url) ToResponse() response.UrlResponse {
	return response.UrlResponse{
		Id:          u.ID,
		OriginalUrl: u.OriginalUrl,
		ShortUrl:    u.ShortUrl,
		View:        u.View,
	}
}

func (u *Url) AddViewUrl() {
	u.View = u.View + 1
}

type UrlRepository interface {
	CreateShortUrl(context.Context, *Url) error
	GetAllUrl(context.Context) ([]Url, error)
	GetUrlByShortUrl(context.Context, string) (Url, error)
	UpdateViewShortUrl(context.Context, Url) error
}

type UrlService interface {
	CreateShortUrl(context.Context, request.UrlRequest) (string, error)
	GetAllUrl(context.Context) ([]response.UrlResponse, error)
	GetUrlByShortUrl(context.Context, string) (response.UrlResponse, error)
}

type UrlController interface {
	CreateShortUrl(*fiber.Ctx) error
	GetAllUrl(*fiber.Ctx) error
	GetUrlByShortUrl(*fiber.Ctx) error
}
