package misc

import (
	"fmt"
	"os"
	"testing"

	"github.com/gocarina/gocsv"
)

type Client struct { // Our example struct, you can use "-" to ignore a field
	Id      string `csv:"不知掉啊"`
	Name    string `csv:"名字呀"`
	Age     string `csv:"年龄呀"`
	NotUsed string `csv:"-"`
}

func TestCsvExport(t *testing.T) {
	clientsFile, err := os.OpenFile("clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer clientsFile.Close()

	clients := []*Client{}

	if err = gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		fmt.Println(err)
	}
	for _, client := range clients {
		fmt.Println("Hello", client.Name)
	}

	if _, err = clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
		fmt.Println(err)
		return
	}

	clients = append(clients, &Client{Id: "12", Name: "你好啊啊", Age: "21"}) // Add clients
	clients = append(clients, &Client{Id: "13", Name: "故事"})
	clients = append(clients, &Client{Id: "14", Name: "水啊", Age: "32"})
	clients = append(clients, &Client{Id: "15", Name: "不知道"})
	csvContent, err := gocsv.MarshalString(&clients) // Get all clients as CSV string
	if err != nil {
		panic(err)
	}
	err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
	fmt.Println(csvContent) // Display all clients as CSV string

}
