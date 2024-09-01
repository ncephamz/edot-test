package main

import (
	"user-service/app/api-user/handler"
	"user-service/business/user"
	"user-service/conf"
	"user-service/pkg/jwt"
	"user-service/pkg/mid"
	"user-service/pkg/password"

	_ "user-service/app/api-user/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v4/pgxpool"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func PrefareRoute(app *fiber.App, cfg conf.Config, db *pgxpool.Pool) {

	// simple common middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Accept, Authorization",
	}))

	// init ratelimiter
	// app.Use(mid.NewLimiter(10, 5))

	// init swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Dependency Injection
	jwtCustom := jwt.Jwt{Secret: cfg.User.SecretKey}
	passCustom := password.Password{}
	secureWithJwt := mid.NewAuthMiddleware(cfg.User.SecretKey)

	// fulfill repo
	userRepo := user.NewRepoUser(db)

	// fullfill usecase / service / core
	userService := user.NewUserService(userRepo, jwtCustom, passCustom)

	// generate handler
	userHandler := handler.NewUserHandler(userService)

	// mapping url
	user := app.Group("/users")
	v1 := user.Group("/v1")
	v1.Post("/register", userHandler.Register)
	v1.Post("/login", userHandler.Login)
	v1.Post("/refresh-token", secureWithJwt, userHandler.RefreshToken)

}
