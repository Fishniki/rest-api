package connection

import (
	"database/sql"
	"log"
	"rest-api/internal/config"
	_ "github.com/lib/pq" 
)

func GetDatabase(conf config.Database) *sql.DB {

	db, err := sql.Open("postgres", conf.URL)
	if err != nil {
		log.Fatal("Failes to open conetion to database", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database", err.Error())
	}  

	return db
}