package main

import (
	"log"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/config"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/handler"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/repository"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/repository/postgres"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/router"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/service"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/storage"
)

func main() {
	if err := config.LoadAllEnv(); err != nil {
		log.Printf("Failed to load config: %v", err)
	}
	cfg, err := config.LoadNewBootCfg()
	if err != nil {
		log.Printf("Failed to load config: %v", err)
	}

	db := storage.InitDBConn(cfg)
	if db != nil {
		defer storage.CloseDBConn(db)
	}

	repos := &repository.Repository{
		Product: postgres.NewProductPostgresRepository(db),
	}
	services := service.NewService(repos, &cfg.CentriConf)
	h := handler.NewHandler(services)
	r := router.NewRouter(h)

	router := r.Init(cfg.TrustedProxies)

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
