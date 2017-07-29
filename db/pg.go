/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/04/11 09:08
 */

package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/wothing/log"
)

var DB *sql.DB

func InitPG(svcName string) {
	//配置docker 数据库连接
	host := os.Getenv("POSTGRES_PORT_5432_TCP_ADDR")
	if host == "" {
		host = "127.0.0.1"
	}
	log.Debug("===================")
	log.Debug(host)
	log.Debug("===================")

	dbhost := GetValue(svcName, "pgsql/host", host)
	dbport := GetValue(svcName, "pgsql/port", "5432")
	dbpwd := GetValue(svcName, "pgsql/password", "")
	dbname := GetValue(svcName, "pgsql/name", "bookcloud")
	dbuser := GetValue(svcName, "pgsql/user", "postgres")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbuser, dbpwd, dbhost, dbport, dbname)
	setupPG(dsn)
}

func setupPG(dbs string, conn ...int) {
	log.Infof(dbs)
	db, err := sql.Open("postgres", dbs)
	if err != nil {
		log.Fatalf("error on connecting to db: %s", err)
	}
	switch len(conn) {
	case 2:
		db.SetMaxOpenConns(conn[0])
		db.SetMaxIdleConns(conn[1])
		db.SetConnMaxLifetime(15 * time.Second)
	default:
		db.SetMaxOpenConns(70)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(10 * time.Second)
	}

	err = db.Ping()
	if err != nil {
		log.Warnf("ping to db: %s", err)
	}

	DB = db
}

func ClosePG() {
	if DB != nil {
		DB.Close()
	}
}
