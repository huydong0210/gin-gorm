package server

import (
	"github.com/gin-gonic/gin"
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

}
