package data

import (
	"time"

	"mangar/backend/internal/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDb() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		time.Sleep(time.Second)
		utils.Warning(err, "Not ready DB")
		return InitDb()
	}
	return db
}

func ConnectDb() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}
