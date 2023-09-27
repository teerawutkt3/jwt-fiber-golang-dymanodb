package routes

import (
	"fiber-poc-api/auth"
	"fiber-poc-api/database/repository"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	jwtWare "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Router(app fiber.Router, db *dynamodb.DynamoDB) {

	middleware := Middleware()
	authRepository := repository.NewUserRepository(db)
	authService := auth.NewAuthService(authRepository)
	authHandler := auth.NewAuthHandler(authService)

	app.Post("/auth/login", authHandler.LoginHandler)
	app.Post("/auth/register", authHandler.RegisterHandler)
	app.Post("/update/user", authHandler.UpdateUserHandler)
	app.Post("/get/user", authHandler.GetUserHandler)
	app.Get("/hello", middleware, hello)
}

func hello(c *fiber.Ctx) error {
	return c.Status(200).JSON("hello test api")
}

func Middleware() fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		SigningKey: jwtWare.SigningKey{Key: []byte(viper.GetString("jwt.secret"))},
	})
}
