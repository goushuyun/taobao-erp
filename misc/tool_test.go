package misc

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/wothing/log"

	"github.com/goushuyun/weixin-golang/pb"

	"github.com/tealeg/xlsx"
)

func TestGenCheckCode(t *testing.T) {
	code := GenCheckCode(4, KC_RAND_KIND_NUM)
	fmt.Println("====>", code)
}

// 获取0-n之间的所有偶数
func even(a int) (array []int) {
	for i := 0; i < a; i++ {
		if i&1 == 0 { // 位操作符&与C语言中使用方式一样
			array = append(array, i)
		}
	}
	return array
}

// 互换两个变量的值
// 不需要使用第三个变量做中间变量
func swap(a, b int) (int, int) {
	a ^= b // 异或等于运算
	b ^= a
	a ^= b
	return a, b
}

// 左移、右移运算
func shifting(a int) int {
	a = a << 1
	a = a >> 1
	return a
}

// 变换符号
func nagation(a int) int {
	// 注意: C语言中是 ~a+1这种方式
	return ^a + 1 // Go语言取反方式和C语言不同，Go语言不支持~符号。
}

func TestBinary(t *testing.T) {
	fmt.Printf("even: %v\n", even(100))
	a, b := swap(100, 200)
	fmt.Printf("swap: %d\t%d\n", a, b)
	fmt.Printf("shifting: %d\n", shifting(100))
	fmt.Printf("nagation: %d\n", nagation(100))
	fmt.Printf("shifting:%d\n", (3 << 1))
	fmt.Printf("shifting:%d\n", ((5 << 1) & 1))
}

func TestNumFormat(t *testing.T) {
	price := 051
	discount := 0.02
	totalFee := float64(price) * discount
	fmt.Println(totalFee)
	totalPriceStr := fmt.Sprintf("%0.0f", totalFee)
	fmt.Println(totalPriceStr)

}
func TestNumFloat(t *testing.T) {
	price := 51
	discountStr := fmt.Sprintf("%.3f", float64(2)/100)
	fmt.Println(discountStr)
	discount := float64(2) / 1000
	fmt.Println(discount)
	totalFee := float64(price) * discount
	fmt.Println(totalFee)
	totalPriceStr := fmt.Sprintf("%0.0f", totalFee)
	fmt.Println(totalPriceStr)

}

func TestSubString(t *testing.T) {
	cardNo := "62284819919020398000"
	str := SubString(cardNo, len(cardNo)-4, 4)
	fmt.Print(str + "\n")
	fmt.Print(len(cardNo))
}

func TestPointerFunc(t *testing.T) {
	var p *int

	test(&p)
	fmt.Println(*p)
}

func test(p **int) {
	x := 100
	*p = &x
	fmt.Println(**p)
}

func TestDownloadAndAnaly(t *testing.T) {

	xlFile, err := xlsx.OpenFile("hello.xlsx")
	if err != nil {
		fmt.Printf("err :%+v", err)
	}
	var books []*pb.Goods
	sheet := xlFile.Sheets[0]
	for _, row := range sheet.Rows {
		isbn, _ := row.Cells[0].String()
		numStr, _ := row.Cells[1].String()
		if isbn == "" {
			break
		}
		book := &pb.Goods{Isbn: isbn, StrNum: numStr}
		books = append(books, book)
	}

	splitList, _ := splitGoodsList(50, books)
	fmt.Printf("%+v\n", len(splitList))
	fmt.Println("======================")
	fmt.Printf("%+v\n", splitList[791])
	fmt.Println("======================")
	fmt.Printf("%+v\n", splitList[792])
	//os.Remove("hello.xls")
}
func TestUrlSubString(t *testing.T) {
	uri := "http://image.goushuyun.cn/Exceltest.xls"
	splitStringArray := strings.Split(uri, "/")
	fmt.Println(splitStringArray)
	fmt.Println(splitStringArray[len(splitStringArray)-1])

	reg := regexp.MustCompile("\\.xlsx$")
	edition := reg.FindString(uri)
	fmt.Println(edition)
	fmt.Println(edition == "")

}

func splitGoodsList(batchSize int, goodsList []*pb.Goods) (splitList [][]*pb.Goods, err error) {

	for i := 0; i < len(goodsList); i += batchSize {

		if i+batchSize >= len(goodsList) {
			splitList = append(splitList, goodsList[i:])
		} else {
			splitList = append(splitList, goodsList[i:i+batchSize])
		}
	}
	return
}

func TestSplitUpload(t *testing.T) {
	var array1 []int = []int{1, 2, 3, 4, 5, 6, 7}
	var array2 []int
	var done bool
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	totalNum := len(array1)
	currentNum := 0
	numChan := make(chan int)
	go func(c chan int, wg *sync.WaitGroup, array1 []int) {
		defer wg.Done()
		defer close(c)
		fmt.Printf("%#v", array1)
		for i := 0; i < len(array1); i++ {
			c <- array1[i]
			numChan <- 1
		}

	}(c, &wg, array1)
	for {
		var v int
		var ok bool
		select {
		case v, ok = <-c:
			if ok {
				fmt.Println("\nchan:", v)
			}
		case v, _ = <-numChan:
			fmt.Println("\nnum:", v)
			currentNum += v
		}

		if currentNum == totalNum {
			break
		}

	}
	wg.Wait()
	done = true
	fmt.Println(done)
	fmt.Println("sfsdfsf", len(array2))
}

func TestWriteExcelAndSave(t *testing.T) {
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
	cell.Value = "I am a cell!"
	// key, err := service.PutExcelFile(file)
	// if err != nil {
	// 	fmt.Print(err)
	//
	// }
	// fmt.Println(key)
	// err = file.Save("MyXLSXFile.xlsx")
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }
}

func TestIsbnInjust(t *testing.T) {
	isbnStr := "9787301176917"
	reg := regexp.MustCompile("^(\\d[- ]*){12}[\\d]$")
	isbn := reg.FindString(isbnStr)
	isbn = strings.Replace(isbn, "-", "", -1)
	isbn = strings.Replace(isbn, " ", "", -1)
	if isbn == "" {
		log.Debug("isbn不正确")

	} else {
		log.Debug("hhh")
	}
}
