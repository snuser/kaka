package dao

import (
	"database/sql"
	"sync"
)

var mysqlInitOnce sync.Once

var dbClient *sql.DB

const MYSQL_DSN = "root:123456@tcp(10.2.0.8:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s"

func init() {
	mysqlInitOnce.Do(func() {
		db, err := sql.Open("mysql", MYSQL_DSN)
		if err != nil {
			panic(err)
		}
		dbClient = db
	})
}

func GetDB() *sql.DB {
	return dbClient
}
