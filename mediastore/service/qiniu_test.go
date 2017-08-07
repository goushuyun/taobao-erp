package service

import (
	"fmt"
	"log"
	"testing"

	"github.com/goushuyun/taobao-erp/pb"
	"github.com/tealeg/xlsx"
)

func TestWriteXLSXFile(t *testing.T) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "My name is Wang Kai !"
	if err != nil {
		fmt.Printf(err.Error())
	}

	// 至此excel file 封装完毕

	PutExcelFile(file)
}

func TestUploadLocal(t *testing.T) {
	err := uploadLocal("../1.jpeg", 0, "test/1.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	t.Log("Maybe it success !!!")
}

func TestFetch(t *testing.T) {

	t.Log("My name is Wang Kai ...")

	url, err := FetchImg(pb.MediaZone_Test, "https://gdp.alicdn.com/imgextra/i1/2356992787/TB2gUWwdrslyKJjSZJiXXb1tFXa_!!2356992787.jpg", "amazon.jpg")

	t.Log(err, url)
}

func TestMakeToken(t *testing.T) {
	token, url := makeToken(0, "kai2.xlsx")

	t.Log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	t.Log(token, url)
	t.Log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
}

func TestGenAt(t *testing.T) {
	t.Log(GenAccessToken("/v2/tune/refresh\n"))
}

func TestRefreshUrls(t *testing.T) {
	err := RefreshURLCache([]string{"http://image.cumpusbox.com/book/9787513557344"})

	if err != nil {
		t.Log(err)
	}
}
