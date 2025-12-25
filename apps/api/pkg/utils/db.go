package utils

import (
	"context"
	"fmt"

	"omnicampus/api/internal/db"
	"omnicampus/api/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() {
	cfg := Config

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
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
