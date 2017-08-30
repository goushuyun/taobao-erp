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
	Id      string `csv:"你好"`
	Name    string `csv:"不知道"`
	Age     string `csv:"你好"`
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

	clients = append(clients, &Client{Id: "12", Name: "Jo事实上hn", Age: "21"}) // Add clients
	clients = append(clients, &Client{Id: "13", Name: "Frs事实上ed"})
	clients = append(clients, &Client{Id: "14", Name: "James", Age: "32"})
	clients = append(clients, &Client{Id: "15", Name: "Danny"})
	err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
}

func TestTaobaoCsvExport(t *testing.T) {

	var taobaoModels = make([]*TaobaoCsvModel, 0, 1)
	describe := "<P><B><font color=blue>基本信息</font></B></P><P>书名:离散数学</P><P>定价：35.00元</P><P>作者:屈婉玲　等编</P><P>出版社：高等教育出版社</P><P>出版日期：2011-01-01</P><P>ISBN：9787040231250</P><P>字数：550000</P><P>页码：</P><P>版次：1</P><P>装帧：平装</P><P>开本：12k</P><P>商品重量：</P><P><B><font color=blue>编辑推荐</font></B></P><HR SIZE=1><P><p>　　本书特色：<br />　　以教育部计算机科学与技术教学指导委员会制订的计算机科学与技术专业规范为指导，内容涵盖计算机科学技术中常用离散结构的数学基础。<br />　　紧密围绕离散数学的基本概念、基本理论精炼选材，体系严谨，内容丰富；面向计算机科学技术，介绍了很多离散数学在计算机科学技术中的应用。<br />　　强化描述与分析离散结构的基本方法与能力的训练，配有丰富的例题和习题；例题有针对性，分析讲解到位；习题难易结合，适合学生课后练习。<br />　　知识体系采用模块化结构，可以根据不同的教学要求进行调整；语言通俗易懂，深入浅出、突出重点、难点，提示易于出错的地方。<br />　　辅助教学资源丰富，配有用于习题课、包含上千道习题的教学辅导用书《离散数学学习指导与习题解析》，PPT电子教案，教学资源库等。</p></P><P><B><font color=blue>内容提要</font></B></P><HR SIZE=1><P><p>　　本书起源于高等教育出版社1998年出版的《离散数学》，是教育部高等学校“九五”规划教材，2004年作为“十五”规划教材出版了修订版。作为“十一五”规划教材，根据教育部计算机科学与技术专业教学指导委员会提出的《计算机科学与技术专业规范》（CCC2005）的教学要求，本教材对内容进行了较多的调整与更新。<br />　　本书分为数理逻辑、集合论、代数结构、组合数学、图论、初等数论等六个部分。全书既有严谨的、系统的理论阐述，也有丰富的、面向计算机科学技术发展的应用实例，同时选配了大量的典型例题与练习。各章内容按照模块化组织，可以适应不同的教学要求。与本书配套的电子教案和习题辅导用书随后将陆续推出。<br />　　本书可以作为普通高等学校计算机科学与技术专业不同方向的本科生的离散数学教材，也可以供其他专业学生和科技人员阅读参考。</p></P><P><B><font color=blue>目录</font></B></P><HR SIZE=1><P></P><P><B><font color=blue>作者介绍</font></B></P><HR SIZE=1><P><p>　　屈婉玲，1969年毕业于北京大学物理系物理专业，现为北京大学信息科学技术学院教授，博士生导师，中国人工智能学会离散数学专委会委员。主要研究方向是算法设计与分析，发表论文20余篇，出版教材、教学参考书、译著20余本，其中包含多本*规划教材和北京市精品教材。所讲授的离散数学课程被评为国家精品课程，两次被评为北京大学十佳教师，并获得北京市教师称号。曾主持过多项国家教材和课程建设项目，并获得北京市教育教学成果（高等教育）一等奖。</p></P><P><B><font color=blue>序言</font></B></P><HR SIZE=1><P></P>"

	model := PackingTaobaoParam("9787040231250", "50000182", "离散数学", "离散数学", "21043943-1_o_2:0:0:;", "北京", "北京", describe, "", "", 2, 10, 7, 8, 9)
	log.Debug(model)
	for i := 0; i < 1; i++ {
		taobaoModels = append(taobaoModels, model)
	}

	filepath := "hello.csv"
	//os.Remove(filepath)
	file, _ := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	//if only file implements the io.Writer interface
	var parser = NewCsv(file)
	err := parser.Parse(taobaoModels)
	if err != nil {
		println(err.Error())
	}

	SetReadOnly(filepath)
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
