/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/07/20 12:07
 */

package db

import (
	"strconv"

	"github.com/jackc/pgx"
	"github.com/wothing/log"
)

var PGX *pgx.ConnPool

func InitPGX(svcName string) {
	dbport := GetValue(svcName, "pgsql/port", "5432")
	port, err := strconv.Atoi(dbport)
	if err != nil {
		port = 5432
	}

	connConfig := pgx.ConnConfig{
		Host:     GetValue(svcName, "pgsql/host", "127.0.0.1"),
		Port:     uint16(port),
		Database: GetValue(svcName, "pgsql/name", "meidb"),
		User:     GetValue(svcName, "pgsql/user", "postgres"),
		Password: GetValue(svcName, "pgsql/password", ""),
	}

	config := pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: 20,
	}
	pool, err := pgx.NewConnPool(config)
	if err != nil {
		log.Warn(err)
	}
	PGX = pool
}

func ClosePGX() error {
	if PGX != nil {
		PGX.Close()
	}
	return nil
}
