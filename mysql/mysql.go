package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/shoother/services/g"
)

func InitMysql(mysqlConfig g.MysqlConfig) *sql.DB {
	db, _ := sql.Open("mysql", mysqlConfig.Addr)
	db.SetMaxOpenConns(mysqlConfig.Max)
	db.SetMaxIdleConns(mysqlConfig.Idle)
	db.Ping()
	return db
}