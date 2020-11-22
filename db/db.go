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
	practiceDB *sql.DB
)

func init() {
	dbPass := os.Getenv("MYSQL_PW")

	var err error
	practiceDB, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}

	err = practiceDB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
