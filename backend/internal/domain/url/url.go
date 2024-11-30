package url

import (
	"backend/internal/domain/url/dto/request"
	"context"

	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	OriginalUrl string
	ShortUrl    string
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

type UrlRepository interface {
	CreateShortUrl(context.Context, *Url) error
}

type UrlService interface {
	CreateShortUrl(context.Context, request.UrlRequest) (string, error)
}

type UrlController interface {
	CreateShortUrl(*fiber.Ctx) error
}
