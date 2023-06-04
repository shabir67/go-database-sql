package godatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go_database")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(30 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
