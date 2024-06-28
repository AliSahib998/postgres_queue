package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"mail-service/internal/configs"
	"mail-service/internal/consumer"
	"mail-service/internal/controller"
	"mail-service/internal/publisher"
	"mail-service/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	appConfigs, err := configs.GetConfigs()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		appConfigs.DB.User,
		appConfigs.DB.Password,
		appConfigs.DB.Host,
		appConfigs.DB.Port,
		appConfigs.DB.DB,
		appConfigs.DB.Schema)
	m, err := migrate.New(
		"file://db/migrations",
		connectionString,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	messagePublisher, err := publisher.NewPGQPublisher(appConfigs.DB)
	if err != nil {
		log.Fatal(err)
	}

	pgqConsumer, err := consumer.NewPGQConsumer(appConfigs.DB, controller.NewController(appConfigs, messagePublisher))
	if err != nil {
		log.Fatal(err)
	}

	err = pgqConsumer.ConsumeMailMessage()
	if err != nil {
		log.Fatal(err)
	}

	appRouter := router.NewRouter(appConfigs, messagePublisher)

	err = appRouter.Run()
	if err != nil {
		log.Fatal(err)
	}

}
