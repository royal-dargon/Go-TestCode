package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "jingzhe:123@tcp/user")
	if err != nil {
		panic(err.Error())
	}
}
