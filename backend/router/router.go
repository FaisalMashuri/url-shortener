package router

import (
	"backend/internal/domain/url"

	"github.com/gofiber/fiber/v2"
)

type RouteParams struct {
	url.UrlController
}

type RouterStruct struct {
	RouteParams RouteParams
}

func NewRouter(params *RouteParams) RouterStruct {
	return RouterStruct{
		RouteParams: *params,
	}
}

func (r *RouterStruct) SetupRoute(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Route("/short-url", func(router fiber.Router) {
		router.Post("/", r.RouteParams.UrlController.CreateShortUrl)
		router.Get("/", r.RouteParams.UrlController.GetAllUrl)
		router.Get("/:shortCode", r.RouteParams.UrlController.GetUrlByShortUrl)
	})
}
