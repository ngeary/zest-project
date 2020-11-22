package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbName   = "zest"
)

var (
	zestDB *sql.DB
)

func init() {
	dbPass := os.Getenv("MYSQL_PW")

	var err error
	zestDB, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}

	err = zestDB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
