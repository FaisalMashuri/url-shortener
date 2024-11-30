package middleware

import (
	"encoding/json"
	"fmt"
	infrastructure "backend/infrastructure/log"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strings"
)

func LogMiddleware(ctx *fiber.Ctx, logString []byte, logger *infrastructure.LogCustom) {
	//var err *fiber.Error
	defer func() {
		if err := recover(); err != nil {
			logger.Logrus.Error("panic err :", err)
		}
	}()
	var requestBody map[string]interface{}
	var responseBody map[string]interface{}
	logInfo := strings.Split(string(logString), "|")
	latency := logInfo[2]
	if ctx.Body() != nil && ctx.Get("Content-Type") == "application/json" {
		if err := json.Unmarshal(ctx.Body(), &requestBody); err != nil {
			requestBody = nil
		}
	}

	if ctx.Response().Body() != nil {
		if err := json.Unmarshal(ctx.Response().Body(), &responseBody); err != nil {
			responseBody = nil
		}
	}

	field := logrus.Fields{
		"endpoint": ctx.Path(),
		"latency":  latency,
		"request_backend": map[string]interface{}{
			"headers": string(ctx.Request().Header.ContentType()),
			"method":  ctx.Method(),
			"body":    requestBody,
		},
		"client_ip": ctx.IP(),
		"response_body": map[string]interface{}{
			"status":  ctx.Response().StatusCode(),
			"headers": string(ctx.Response().Header.ContentType()),
			"body":    responseBody,
		},
		"request_id": ctx.Locals("request-id"),
	}

	if len(logInfo) >= 7 {
		Error(ctx, logger, latency)
	} else {
		logger.Logrus.WithFields(field).Info()
	}
}

func Error(ctx *fiber.Ctx, logger *infrastructure.LogCustom, latency string) {
	var requestBody map[string]interface{}
	var responseBody map[string]interface{}
	var errStr, rootCause, errorCode string
	var errorCause []string
	if ctx.Body() != nil && ctx.Get("Content-Type") == "application/json" {
		if err := json.Unmarshal(ctx.Body(), &requestBody); err != nil {
			requestBody = nil
		}
	}

	if ctx.Response().Body() != nil {
		if err := json.Unmarshal(ctx.Response().Body(), &responseBody); err != nil {
			responseBody = nil
		}
	}
	if ctx.Locals("error") != nil {
		errStr = ctx.Locals("error").(string)
		errorCause = strings.Split(errStr, "\n")
		rootCause = strings.Replace(errorCause[2], "\t", "", 1)
		errorCode = errorCause[0]
	} else {
		errStr = ""
		rootCause = ""
		errorCode = ""
	}
	fmt.Println(errorCode)

	field := logrus.Fields{
		"endpoint": ctx.Path(),
		"latency":  latency,
		"request_backend": map[string]interface{}{
			"headers": string(ctx.Request().Header.ContentType()),
			"method":  ctx.Method(),
			"body":    requestBody,
		},
		"client_ip": ctx.IP(),
		"response_body": map[string]interface{}{
			"status":  ctx.Response().StatusCode(),
			"headers": string(ctx.Response().Header.ContentType()),
			"body":    responseBody,
		},
		"request_id":    ctx.Locals("request-id"),
		"package_name":  ctx.Get("pkg_name"),
		"error_messgae": errorCode,
		"error_cause":   rootCause,
	}
	logger.Logrus.WithFields(field).Error()
}
