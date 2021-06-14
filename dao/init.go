package dao

import (
	"database/sql"
	"fmt"
	"sync"
)

var mysqlInitOnce sync.Once

var dbClient *sql.DB

const MYSQL_DSN = "root:123456@tcp(10.2.0.8:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s"

func init() {
	mysqlInitOnce.Do(func() {
		db, err := sql.Open("mysql", MYSQL_DSN)
		if err != nil {
			panic(fmt.Sprintf("mysql 初始化失败 err:%s\n", err))
		}
		//先注释， 保证没有数据库连接的情况下能够正常使用
		//err = db.Ping()
		//if err != nil {
		//	panic(fmt.Sprintf("mysql 初始化失败 err:%s\n", err))
		//}
		dbClient = db
	})
}

func GetDB() *sql.DB {
	return dbClient
}
