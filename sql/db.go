package sql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

var pool *pgxpool.Pool

func Init() {
	var err error

	if pool == nil {
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		name := os.Getenv("POSTGRES_DB")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")

		connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, name)
		pool, err = pgxpool.Connect(context.Background(), connectionString)

		if err != nil {
			panic("Cannot connect to DB")
		}
	}
}
