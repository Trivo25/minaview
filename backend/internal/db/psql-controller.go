package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func GetCon() *pgx.Conn {
	// TODO: edit conn string
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:Florian14@127.0.0.1:5432/minaview_dev")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	/*	defer conn.Close(context.Background())
	 */
	return conn
}
