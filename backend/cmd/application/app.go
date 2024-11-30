package application

import (
	"backend/config"
	"backend/infrastructure/database"
	infrastructure "backend/infrastructure/log"
	redis_client "backend/infrastructure/redis"
	urlController "backend/internal/domain/url/controller"
	urlRepository "backend/internal/domain/url/repository"
	urlService "backend/internal/domain/url/service"
	userController "backend/internal/domain/user/controller"
	userRepository "backend/internal/domain/user/repository"
	userService "backend/internal/domain/user/service"
	middleware "backend/middleware/error"
	middlewareLog "backend/middleware/log"
	"backend/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

func Run() {
	config.LoadConfig()
	log := infrastructure.NewLogCustom()
	err := middleware.LoadErrorListFromJsonFile(config.AppConfig.ErrorContract.JSONPathFile)
	if err != nil {
		log.Logrus.Fatal("Failed to read to errorContract.json:", err)
	}

	// Database
	db := database.ConnectDatabase()

	// redis
	redisClient := redis_client.RedisClient

	//// Kafka
	//_ = kafka.NewKafkaConsumer(*log)
	//_ = kafka.NewKafkaProducer(*log)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Jakarta",
		Done: func(ctx *fiber.Ctx, logString []byte) {
			middlewareLog.LogMiddleware(ctx, logString, log)
		},
	}))

	app.Use(requestid.New(requestid.Config{
		Generator:  utils.UUIDv4,
		ContextKey: "request-id",
	}))

	//Todo : Define Repository here
	redisRepo := redis_client.NewRedisRepository(redisClient)
	userRepo := userRepository.NewRepository()
	urlRepo := urlRepository.NewUrlRepository(db.DB)

	//Todo : Define Service here
	userSvc := userService.NewService(db.DB, userRepo, &redisRepo)
	urlSvc := urlService.NewUrlService(urlRepo, redisRepo)

	//Todo: Define controller
	userCtrl := userController.NewController(userSvc, log)
	urlCtrl := urlController.NewUrlController(urlSvc)

	routerApp := router.NewRouter(&router.RouteParams{
		userCtrl,
		urlCtrl,
	})
	routerApp.SetupRoute(app)
	err = app.Listen(fmt.Sprintf(":%s", config.AppConfig.AppConfig.Port))
	if err != nil {
		log.Logrus.Fatal("Failed to start server:", err)
	}
}
