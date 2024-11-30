package router

import (
	"backend/internal/domain/url"
	"backend/internal/domain/user"

	"github.com/gofiber/fiber/v2"
)

type RouteParams struct {
	user.UserController
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
	v1.Route("/auth", func(router fiber.Router) {
		router.Post("/login", r.RouteParams.UserController.Login)
	})

	v1.Route("/short-url", func(router fiber.Router) {
		router.Post("/", r.RouteParams.CreateShortUrl)
	})
}