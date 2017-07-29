package misc

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/goushuyun/weixin-golang/pb"
)

const (
	filename = "/var/log/bookcloud/record.log"
)

func LogErrOrder(order *pb.Order, impact string, err error) {
	line := "time:%#v, order:%s Error occurred ,store's id is %s,user's id is %s and trade channel is %s ,trade no is %s ,the impact is %s !!!. err reason is %#v \n"
	wireteString := fmt.Sprintf(line, time.Now(), order.Id, order.StoreId, order.UserId, order.PayChannel, order.TradeNo, impact, err)
	var file *os.File

	if checkFileIsExist(filename) { //如果文件存在
		file, _ = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		file, _ = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}

	io.WriteString(file, wireteString) //写入文件(字符串)

}
func LogErrAccount(account *pb.AccountItem, impact string, err error) {
	line := "time:%#v, account insert Error occurred ,order'id is %s,store's id is %s ,the impact is %s !!!. err reason is %#v \n"
	wireteString := fmt.Sprintf(line, time.Now(), account.OrderId, account.StoreId, impact, err)
	var file *os.File

	if checkFileIsExist(filename) { //如果文件存在
		file, _ = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		file, _ = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}

	io.WriteString(file, wireteString) //写入文件(字符串)

}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
