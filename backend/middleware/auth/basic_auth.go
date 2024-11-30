package auth

import (
	"backend/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/pkg/errors"
)

func NewBasicAuthMiddleware() basicauth.Config {
	return basicauth.Config{
		Users: map[string]string{
			"username": "Gojo",
			"password": "Satoru",
		},
		Realm: "Forbidden",
		Authorizer: func(username, password string) bool {
			if username == "Gojo" && password == "Satoru" {
				return true
			}
			return false
		},
		Unauthorized: func(ctx *fiber.Ctx) error {
			return errors.New(shared.Unauthorized)
		},
		ContextPassword: "_pass",
		ContextUsername: "_user",
	}
}
