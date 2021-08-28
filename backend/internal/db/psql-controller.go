package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

func GetCon() *pgxpool.Pool {
	// TODO: edit conn string
	conn, err := pgxpool.Connect(context.Background(), "postgres://postgres:Florian14@127.0.0.1:5432/minaview_dev")
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		if err != nil {
			return nil
		}
		os.Exit(1)
	}
	/*	defer conn.Close(context.Background())
	 */
	return conn
}
