package main

import (
	"github.com/trjade/wallet-transfer-service/internal/config"
	"github.com/trjade/wallet-transfer-service/internal/db"
	"github.com/trjade/wallet-transfer-service/internal/logger"
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

	// Initialize and start the server here
	log.Info("Server is starting...")

}
