package db

import (
	cfg "cloudisk/config"
	"database/sql"
	"log"
)

type Mysql struct {
	DbName   string
	Address  string
	Password string
}

var db *sql.DB

func InitConnect() {
	var err error
	db, err = sql.Open("mysql", cfg.MysqlSource)
	if err != nil {
		log.Print("mysql open failed")
		return
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		log.Print("db connect failed")
		return
	}
}

func DBConn() *sql.DB {
	return db
}
