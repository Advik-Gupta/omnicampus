package utils

import (
	"context"
	"fmt"

	"omnicampus/api/internal/config"
	"omnicampus/api/internal/db"
	"omnicampus/api/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() {
	cfg := config.Load()

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		panic(err)
	}

	db.Queries = sqlc.New(pool)

	if db.Queries == nil {
		panic("db.Queries is nil after InitDB")
	}

	fmt.Println("Postgres connected successfully")
}
