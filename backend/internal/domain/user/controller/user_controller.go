package controller

import (
	"fmt"
	infrastructure "backend/infrastructure/log"
	"backend/internal/domain/user"
	"backend/internal/domain/user/dto/requests"
	middleware "backend/middleware/error"
	validation "backend/middleware/validate"
	"backend/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"reflect"
)

type userController struct {
	service user.UserService
	logger  *infrastructure.LogCustom
}

func NewController(service user.UserService, logger *infrastructure.LogCustom) user.UserController {
	return &userController{
		service: service,
		logger:  logger,
	}
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	var loginReq requests.LoginRequest
	err := ctx.BodyParser(&loginReq)
	if err != nil {
		ctx.Locals("error", fmt.Sprintf("%+v", errors.Cause(errors.WithStack(err))))
		ctx.Locals("pkg_name", reflect.TypeOf(userController{}).PkgPath())
		fmt.Println("ERROR PARSING : ", err)
		return errors.New(shared.ErrInvalidRequestFamily)
	}
	fieldErr := validation.ValidateRequest(loginReq)
	if fieldErr != nil {
		fmt.Println("ERROR : ", fieldErr)
		return errors.New(fieldErr.Error())
	}
	res, err := c.service.Login(ctx.Context(), &loginReq)
	if err != nil {

		return errors.New(shared.ErrUnexpectedError)
	}
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, res)
}

func (c *userController) Register(ctx *fiber.Ctx) error {
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)

}

func (c *userController) GetDetailUserJWT(ctx *fiber.Ctx) error {
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)
}

func (c *userController) VerifyUser(ctx *fiber.Ctx) error {
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)
}

func (c *userController) ResendOTP(ctx *fiber.Ctx) error {
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)
}
