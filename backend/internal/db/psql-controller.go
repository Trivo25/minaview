package db

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"os"
)

func GetCon() (*pgxpool.Pool, error) {

	err := godotenv.Load()

	if err != nil {
		return nil, errors.New("An Error has occured")
	}

	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, errors.New("An Error has occured")
	}
	return conn, nil
}
