package driver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Conn struct {
	*sql.DB
}

func NewDBConn() Conn {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		log.Fatal("Database connect error",err)
	}

	conn := Conn{db}

	return conn
}