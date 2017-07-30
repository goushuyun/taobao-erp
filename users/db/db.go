package db

import (
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"

	. "github.com/goushuyun/taobao-erp/db"
)

func SaveUser(u *pb.User) error {
	query := "insert into users (mobile, password, name) values($1, $2, $3) returning id"

	log.Debugf("insert into users (mobile, password, name) values(%s, %s, %s) returning id", u.Mobile, u.Password, u.Name)
	err := DB.QueryRow(query, u.Mobile, u.Password, u.Name).Scan(&u.Id)

	return err
}
