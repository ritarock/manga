package data

const (
	DBMS     = "mysql"
	PROTOCOL = "tcp(db:3306)"
	USER     = "user"
	PASS     = "password"
	DBNAME   = "app"
	CONNECT  = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
)
