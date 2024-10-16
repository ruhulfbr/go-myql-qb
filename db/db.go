package db

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect(username, password, host, dbname string) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname)

	var DBConnection *sql.DB
	var err error
	DBConnection, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error creating the database DBConnection: %v", err)
	}

	if err := DBConnection.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	return DBConnection
}

func Close(DBConnection *sql.DB) {
	if DBConnection != nil {
		err := DBConnection.Close()
		if err != nil {
			log.Fatalf("Error closing the database DBConnection: %v", err)
		}
	}
}

func IsConnected(DBConnection *sql.DB) bool {
	if DBConnection == nil {
		log.Fatalf("Database is not connected.")
	}

	return true
}
