package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ruhulfbr/go-mysql-qb/builder"
	"github.com/ruhulfbr/go-mysql-qb/db"
)

var DBConnection *sql.DB

func ConnectDB(username, password, host, dbname string) {
	DBConnection = db.Connect(username, password, host, dbname)
}

func CloseDB() {
	db.Close(DBConnection)
}

func Table(table string) *builder.QueryBuilder {
	return builder.Table(DBConnection, table)
}
