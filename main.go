package main

import (
	"fiber-poc-api/routes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"strings"
)

func main() {
	// ==> Get config from config.yaml
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("fatal error config file: %+v \n", err)
	}

	// ==> Connect database mysql
	db := databaseConnection()
	fmt.Sprintf("sdf %+v", db)

	app := fiber.New()
	// ==> cors
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// ==> routes
	api := app.Group(viper.GetString("server.context-path"))
	routes.Router(api, db)

	// ==> server start
	port := viper.GetString("server.port")
	app.Listen(fmt.Sprintf(":%s", port))
}

func databaseConnection() *dynamodb.DynamoDB {
	config := &aws.Config{
		Region:      aws.String(viper.GetString("database.region")),
		//Endpoint:    aws.String(viper.GetString("database.url")),
		Credentials: credentials.NewStaticCredentials(viper.GetString("database.key"), viper.GetString("database.secret"), ""),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		return nil
	}
	svc := dynamodb.New(sess)
	return svc
}
