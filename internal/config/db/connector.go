package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetDB() (db *sql.DB) {
	ps := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		`localhost`, `user`, `user`, `postgre`)
	db, err := sql.Open("pgx", ps)
	if err != nil {
		panic(err)
	}
	return db
}
