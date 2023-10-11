package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"shopstoretest/cfg"
	"time"
)

type MySQLDB struct {
	DB *sql.DB
}

func New(cfg cfg.DataBaseCfg) *MySQLDB {
	//fmt.Printf("%s:%s@%s(%s:%s)/%s",
	//cfg.UserName, cfg.Password, cfg.Protocol, cfg.Host, cfg.Port, cfg.Database)

	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
	//	cfg.UserName, cfg.Password, cfg.Protocol, cfg.Host, cfg.Port, cfg.Database))

	db, err := sql.Open("mysql", "shopstore:shopstoret0lk2o20@(localhost:7777)/shopstore_db")

	//fmt.Println("hellloooooo", err)
	if err != nil {

		fmt.Println("fuck this error")
		panic(fmt.Errorf("cant open mysql db %v", err))
	}
	fmt.Println("im here")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{DB: db}
}
