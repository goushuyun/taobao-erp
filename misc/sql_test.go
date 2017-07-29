package misc

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/goushuyun/weixin-golang/db"
	"github.com/goushuyun/weixin-golang/pb"
	"github.com/wothing/log"
)

func TestPgArray2StringSlice(t *testing.T) {
	cases := []struct {
		str      []byte
		expected []string
	}{
		{[]byte("{abc,de,f}"), []string{"abc", "de", "f"}},
		{[]byte("{abc}"), []string{"abc"}},
		{[]byte("{abc,}"), []string{"abc"}},
		{[]byte("{abc} "), []string{"abc"}},
		{[]byte("{abc"), []string{"abc"}},
		{[]byte("{}"), []string{}},
		{[]byte("}"), []string{}},
		{[]byte("{"), []string{}},
		{[]byte("x"), []string{"x"}},
		{nil, []string{}},
	}
	for i, c := range cases {
		actual := PgArray2StringSlice(c.str)
		if !reflect.DeepEqual(c.expected, actual) {
			t.Fatalf("case %d expected string array %#v, but got %#v", i, c.expected, actual)
		}
	}
}

func TestTransaction(t *testing.T) {
	db.InitPG("hello")
	defer db.ClosePG()
	for i := 0; i < 10; i++ {
		go update1("mobile"+strconv.Itoa(i), "nickname"+strconv.Itoa(i), "00000001")
	}
	time.Sleep(15 * time.Second)
}

func update1(mobile, nickname, id string) {
	tx, _ := db.DB.Begin()

	var status int

	_, err := tx.Exec("update seller set mobile=$1 ,nickname=$2,status=status+1 where id=$3", mobile, nickname, id)

	if err != nil {
		log.Debugf("执行出错:%+v", err)

		return
	}
	log.Debugf("查询开始")
	err = db.DB.QueryRow("select status from seller where id=$1", id).Scan(&status)
	if err != nil {
		log.Debugf("查询出错:%+v", err)
		return
	}
	log.Debugf("查询结束")
	log.Debugf("开始状态，%d", status)
	if status > 30 {
		log.Debugf("事务回滚")
		err = tx.Rollback()
		if err != nil {
			log.Debugf("提交出错%+v", err)
		}
		return
	}
	log.Debugf("事务提交开始")
	err = tx.Commit()
	log.Debugf("事务提交结束")
	if err != nil {
		log.Debugf("提交出错%+v", err)
		return
	}
}
func TestString(t *testing.T) {
	var (
		args []interface{}
	)

	ids := strings.FieldsFunc("5f1c7395-c45a-481f-b342-9acff50b1da2,9c21c854-d0d3-4586-9331-dfe087377d9c", split)

	if len(ids) > 0 {
		stmt := `and category in (${ids})  `
		stmt = strings.Replace(stmt, "${"+"ids"+"}",
			strings.Repeat(",'$%s'", len(ids))[1:], -1)
		for _, s := range ids {
			args = append(args, s)
		}
		condition := fmt.Sprintf(stmt, args...)
		fmt.Print(condition)
	}

}
func split(s rune) bool {
	if s == ',' {
		return true
	}
	return false
}

func TestJsonb(t *testing.T) {
	db.InitPG("hello")
	defer db.ClosePG()
	query := "select after_sale_images from orders where id='17042700000027'"
	var images []*pb.AfterSaleImage
	var imageStr string
	err := db.DB.QueryRow(query).Scan(&imageStr)
	if err != nil {
		log.Debug(err)
		return
	}
	if err := json.Unmarshal([]byte(imageStr), &images); err == nil {
		fmt.Println("================json str 转struct==")
	}

	for i := 0; i < len(images); i++ {
		fmt.Println("==============================")
		fmt.Print(images[i].Url)
		fmt.Println("==============================")
	}
	json.Marshal(imageStr)
	log.Debug("===============")
	log.Debugf("======%#v", images)
	log.Debug("===============")

	if b, err := json.Marshal(images); err == nil {
		fmt.Println("================struct 到json str==")
		fmt.Println(string(b))
	}
	return
}
func TestRegPay(t *testing.T) {
	payAlipayWeb := "alipay_wap"
	payAlipay := "alipay"
	payWx := "wx"
	payWxWeb := "wx_web"

	fmt.Println(strings.Contains(payAlipayWeb, "lipay"))
	fmt.Println(strings.Contains(payAlipay, "alipay"))
	fmt.Println(strings.Contains(payAlipay, "wx"))
	fmt.Println(strings.Contains(payAlipayWeb, "wx"))
	fmt.Println(strings.Contains(payWx, "alipay"))
	fmt.Println(strings.Contains(payWxWeb, "alipay"))
	fmt.Println(strings.Contains(payWx, "wx"))
	fmt.Println(strings.Contains(payWxWeb, "wx"))
}

func TestTimeUnix(t *testing.T) {
	now := time.Now()
	statisticDate := now.Add(-1 * 24 * time.Hour)
	fmt.Println(statisticDate.Format("2006-01-02"))
	tm := time.Unix(time.Now().Unix(), 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
	fmt.Print(now.Unix())
	startAt := (now.Add(-1 * 24 * time.Hour)).Format("2006-01-02")
	endAt := (now.Add(-15 * 24 * time.Hour)).Format("2006-01-02")
	fmt.Printf("\n start_at:%s and end_at:%s", startAt, endAt)
	fmt.Print("=========月份计算==========\n")
	fmt.Println(tm.Format("2006-01"))
	fmt.Println(tm.AddDate(0, -1, 0))
	fmt.Println(tm.AddDate(0, -2, 0))
	fmt.Print("===================\n")
}
