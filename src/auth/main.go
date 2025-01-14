package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/qmessentials/qme-enterprise/auth/repositories"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	config, err := lib.NewConfig()
	if err != nil {
		logger.Error("Unable to read configuration", "error", err)
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=America/New_York",
		config.Postgres.DBHost, config.Postgres.DBUser, config.Postgres.DBPassword,
		config.Postgres.DBName, config.Postgres.DBPort)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		logger.Error("Unable to connect to database", "error", err)
		panic(err)
	}
	defer conn.Close(context.Background())
	schemaVersionRepo := repositories.NewSchemaVersionRepository(conn)
	schemaVersion, err := schemaVersionRepo.GetScalar()
	if err != nil {
		logger.Error("Database test query failed", "error", err)
		panic(err)
	}
	logger.Info("Database test query succeeded", "Schema Version", schemaVersion)

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK", "timestamp": time.Now()})
	})
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
