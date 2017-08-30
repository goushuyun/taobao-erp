package export

import (
	"fmt"
	"os"
	"testing"

	"github.com/gocarina/gocsv"
	"github.com/goushuyun/taobao-erp/misc"
	"github.com/wothing/log"
)

type People struct {
	Name   string `csv:"姓名"`
	Age    int
	Sex    string
	Weight float32 `csv:"体重"`
	Marry  bool    `csv:"婚姻状况"`
}
type Client struct { // Our example struct, you can use "-" to ignore a field
	Id      string `csv:"client_id"`
	Name    string `csv:"client_name"`
	Age     string `csv:"client_age"`
	NotUsed string `csv:"-"`
}

func TestNew(t *testing.T) {
	clientsFile, err := os.OpenFile("clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []*Client{}

	for _, client := range clients {
		fmt.Println("Hello", client.Name)
	}

	if _, err = clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	clients = append(clients, &Client{Id: "12", Name: "John", Age: "21"}) // Add clients
	clients = append(clients, &Client{Id: "13", Name: "Fred"})
	clients = append(clients, &Client{Id: "14", Name: "James", Age: "32"})
	clients = append(clients, &Client{Id: "15", Name: "Danny"})
	err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
}

func TestCsvExport(t *testing.T) {

	var peoples = make([]*People, 0, 10)
	for i := 0; i < 10; i++ {
		var people = &People{
			Name:   "李肖",
			Age:    22,
			Sex:    "帅哥",
			Weight: 130.6,
			Marry:  false,
		}
		peoples = append(peoples, people)
	}

	file, _ := os.OpenFile("aa.csv", os.O_CREATE|os.O_RDWR, 0644)

	//if only file implements the io.Writer interface
	var parser = NewCsv(file)
	err := parser.Parse(peoples)
	if err != nil {
		println(err.Error())
	}
	file.Close()

}

func TestTaobaoCsvExport(t *testing.T) {

	// var taobaoModels = make([]*TaobaoCsvModel, 0, 1)
	// model := PackingTaobaoParam("9787040231250", "50000182", "离散数学", "离散数学", "21043943-1_o_2:0:0:;", "北京", "北京", describe, 2, 10, 7, 8, 9)
	// log.Debug(model)
	// taobaoModels = append(taobaoModels, model)
	// filepath := "hello.csv"
	// os.Remove(filepath)
	// file, _ := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	// defer file.Close()
	// //if only file implements the io.Writer interface
	// var parser = NewCsv(file)
	// err := parser.Parse(taobaoModels)
	// if err != nil {
	// 	println(err.Error())
	// }
	//
	// SetReadOnly(filepath)
	return

}

func TestDownloadFile(t *testing.T) {

	rawURL := "http://taoimage.goushuyun.cn/201708110637139787300056074.jpg"
	err := misc.DownloadFileFromServer("test.tbi", rawURL)
	if err != nil {
		log.Error(err)
		return
	}
}
