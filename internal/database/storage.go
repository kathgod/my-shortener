package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbOpenError = "Open DataBase Error"
)

// OpenDB Функция открытия базы данных.
func OpenDB(DBDSN string) *sql.DB {
	db, errDB := sql.Open("postgres", DBDSN)
	log.Println(DBDSN)
	if errDB != nil {
		log.Println(dbOpenError)
		log.Println(errDB)
	}
	return db
}
