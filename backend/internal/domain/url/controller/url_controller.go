package controller

import (
	"backend/internal/domain/url"
	"backend/internal/domain/url/dto/request"
	middleware "backend/middleware/error"
	"backend/shared"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type urlController struct {
	service url.UrlService
}

// CreateShortUrl implements url.UrlController.
func (controller *urlController) CreateShortUrl(ctx *fiber.Ctx) error {
	var reqData request.UrlRequest
	err := ctx.BodyParser(&reqData)
	if err != nil {
		fmt.Println("ERROR FAMILY : ", err)
		ctx.Locals("error", fmt.Sprintf("%+v", errors.Cause(errors.WithStack(err))))
		ctx.Locals("pkg_name", reflect.TypeOf(urlController{}).PkgPath())
		fmt.Println("ERROR PARSING : ", err)
		return errors.New(shared.ErrInvalidRequestFamily)
	}
	data, err := controller.service.CreateShortUrl(ctx.Context(), reqData)
	if err != nil {
		fmt.Println("error : ", err)
		return errors.New(shared.ErrUnexpectedError)
	}
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, data)
}

func NewUrlController(service url.UrlService) url.UrlController {
	return &urlController{
		service: service,
	}
}
