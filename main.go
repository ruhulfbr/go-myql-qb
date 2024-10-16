package DB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ruhulfbr/go-mysql-qb/builder"
	"github.com/ruhulfbr/go-mysql-qb/db"
)

var Connection *sql.DB

func ConnectDB(username, password, host, dbname string) {
	Connection = db.Connect(username, password, host, dbname)
}

func CloseDB() {
	db.Close(Connection)
}

func Table(table string) *builder.QueryBuilder {
	return builder.Table(Connection, table)
}
