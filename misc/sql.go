// Copyright 2015-2016, Wothing Co., Ltd.
// All rights reserved.

package misc

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	"github.com/pborman/uuid"
	"github.com/wothing/log"
)

func StringSlice2PgArray(arr []string) string {
	if len(arr) == 0 {
		return "{}"
	}

	return "{" + strings.Join(arr, ",") + "}"
}

func PgArray2StringSlice(bs []byte) []string {
	s := strings.Trim(string(bs), "{}, ")
	if s == "" {
		// because strings.Split("", ",") return []string{""}
		return []string{}
	}
	return strings.Split(s, ",")
}

func PgArray2IntSlice(bs []byte) []int {
	s := strings.Trim(string(bs), "{}")
	res := []int{}

	for _, v := range strings.Split(s, ",") {
		num, _ := strconv.Atoi(v)
		res = append(res, num)
	}

	return res
}

func GenIdListSql(idList []string) (string, error) {
	var sql string
	if len(idList) == 0 {
		return "", errors.New("idList lenth is 0")
	}

	for i, uid := range idList {
		tmp := uuid.Parse(uid)
		if tmp == nil {
			return "", errors.New("id[" + strconv.Itoa(i) + "] is not uuid")
		}
	}

	for i, id := range idList {
		if i == 0 {
			sql = "('" + id
			continue
		}
		sql = sql + "', '" + id
	}

	sql = sql + "')"
	return sql, nil
}

func RollbackCtx(ctx context.Context, t *sql.Tx) {
	err := t.Rollback()
	if err != nil && err != sql.ErrTxDone {
		log.Terrorf(GetTidFromContext(ctx), "Roll back fail, err:", err.Error())
	}
}

func Rollback(tid string, t *sql.Tx) {
	err := t.Rollback()
	if err != nil && err != sql.ErrTxDone {
		log.Terrorf(tid, "Roll back fail, err:", err.Error())
	}
}

//check if it is already exist
func IsDuplicate(err error) bool {
	return strings.Contains(err.Error(), "duplicate")
}

//check not found
func NotFound(err error) bool {
	return strings.Contains(err.Error(), "syntax for uuid")
}

func GenArraySqlString(args []string) string {
	if len(args) == 0 {
		return ""
	}

	sql := ""
	// select id from users where id in ('uuid1', 'uuid2', 'uuid3'
	for i, id := range args {
		if i == 0 {
			sql = sql + `"` + id + `"`
			continue
		}
		sql = sql + "," + `"` + id + `"`
	}
	return sql
}

func ByteArray2String(b []byte) []string {
	s := string(b)
	r := strings.Trim(s, "{}")
	var a []string
	for _, t := range strings.Split(r, ",") {
		a = append(a, t)
	}
	return a
}

func Enums2PGArrayWithoutBracket(x interface{}) string {
	s := fmt.Sprintf("%d", x)
	s = strings.Trim(s, "[]")
	return strings.Replace(s, " ", ",", -1)
}

func Enums2PGArray(x interface{}) string {
	return "{" + Enums2PGArrayWithoutBracket(x) + "}"
}

// Mostly used for gen SQL IN
// ex:
// a:=[]string{"uuid1","uuid2"}
// in := Generate(a,"{","}","'",",")
// will result: {'uuid1','uuid2'}
func Generate(x interface{}, prefix, suffix, wrapper, separator string) string {
	str := fmt.Sprint(x)
	str = strings.Trim(str, `[]`)

	if wrapper != "" {
		array := strings.Split(str, " ")
		for i, v := range array {
			array[i] = wrapper + v + wrapper
		}

		str = strings.Join(array, separator)
	}

	str = strings.Replace(str, " ", separator, -1)

	return prefix + str + suffix
}
