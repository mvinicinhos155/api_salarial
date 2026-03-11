package database


import (
	"database/sql"
	"os"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {

	connStr := os.Getenv("DATABASE_URL")

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Banco conectado com sucesso")

	return db, nil
}