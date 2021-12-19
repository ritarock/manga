package data

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DBMS     = "mysql"
	PROTOCOL = "tcp(db:3306)"
	USER     = "user"
	PASS     = "pass"
	DBNAME   = "app"
	CONNECT  = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
)

func InitDb() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		time.Sleep(time.Second)
		fmt.Println(err, "Not ready DB")
		return InitDb()
	}
	return db
}

func connectDb() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}
