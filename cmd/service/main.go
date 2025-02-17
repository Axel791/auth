package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Axel791/auth/internal/config"
	"github.com/Axel791/auth/internal/db"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(logrus.InfoLevel)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	dbConn, err := db.ConnectDB(cfg.DatabaseDSN, cfg)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer func() {
		if dbConn != nil {
			_ = dbConn.Close()
		}
	}()

	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Logger)
}
