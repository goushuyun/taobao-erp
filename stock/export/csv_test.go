package export

import (
	"os"
	"testing"
)

type People struct {
	Name   string `csv:"姓名"`
	Age    int
	Sex    string
	Weight float32 `csv:"体重"`
	Marry  bool    `csv:"婚姻状况"`
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
