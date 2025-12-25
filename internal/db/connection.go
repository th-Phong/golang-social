package db

import (
	"context"
	"fmt"
	"log"
	"phongtran/go-social/golang-social/internal/config"
	"phongtran/go-social/golang-social/internal/db/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB   sqlc.Querier
	Pool *pgxpool.Pool
)

func InitDB() error {
	dns := config.NewConfig().DNS()

	conf, err := pgxpool.ParseConfig(dns)
	if err != nil {
		return fmt.Errorf("parse config error: %v", err)
	}

	conf.MaxConns = 50
	conf.MinConns = 5
	conf.MaxConnLifetime = 30 * time.Minute
	conf.MaxConnIdleTime = 5 * time.Minute
	conf.HealthCheckPeriod = 1 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Pool, err = pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return fmt.Errorf("init DB pool error: %v", err)
	}

	DB = sqlc.New(Pool)

	if err := Pool.Ping(ctx); err != nil {
		return fmt.Errorf("ping DB error: %v", err)
	}

	log.Println("Database connection established")

	return nil
}
