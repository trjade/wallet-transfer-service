package main

import (
	"github.com/trjade/wallet-transfer-service/internal/config"
	"github.com/trjade/wallet-transfer-service/internal/db"
	"github.com/trjade/wallet-transfer-service/internal/logger"
	"go.uber.org/zap"
)

func main() {
	cnfg := config.Load()

	log, err := logger.New()
	if err != nil {
		panic(err)
	}

	dbPool, err := db.NewPostgresDB(cnfg.DBURL)
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	if err := db.RunMigrations(cnfg.DBURL); err != nil {
		log.Fatal("Failed to run database migrations", zap.Error(err))
	}
	log.Info("database migrations applied")

	// Initialize and start the server here
	log.Info("Server is starting...")

}
