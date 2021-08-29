package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"os"
)

func GetCon() (*pgxpool.Pool, error) {

	err := godotenv.Load()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	/*	defer conn.Close(context.Background())
	 */
	return conn, nil
}
