package main

import (
	"github.com/gin-gonic/gin"
	"todo-list-gin-gorm/internal/api"
	"todo-list-gin-gorm/internal/config"
	"todo-list-gin-gorm/internal/database"
	"todo-list-gin-gorm/pkg/logger"
)

func main() {
	log := logger.NewLogger()

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal().Err(err).Msg("load config file failed")
	}

	db, err := database.Initialize(cfg.DatabaseUrl)

	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize database: %s", cfg.DatabaseUrl)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatal().Err(err).Msgf("Failed to migrate database: %s", cfg.DatabaseUrl)
	}

	router := gin.Default()
	api.SetUpRoutes(router, db)

	err = router.Run(cfg.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

}
