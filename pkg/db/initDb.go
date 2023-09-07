package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/user")
	if err != nil {
		return nil, err
	}
	return db, nil
}