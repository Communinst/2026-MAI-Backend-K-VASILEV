package main

import (
	"log"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/internal/config"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/internal/handler"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/internal/router"
)

func main() {
	if err := config.LoadAllEnv(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	cfg, err := config.LoadNewBootCfg()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	h := handler.NewHandler()
	r := router.NewRouter(h)

	router := r.Init(cfg.TrustedProxies)

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
