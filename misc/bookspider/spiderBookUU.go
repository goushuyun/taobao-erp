package bookspider

import (
	"fmt"
	"regexp"
	"strings"

	log "github.com/wothing/log"

	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
)

type BookUUListProcesser struct {
}

func NewBookUUListProcesser() *BookUUListProcesser {
	return &BookUUListProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (s *BookUUListProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		log.Debug(p.Errormsg())
		return
	}
	query := p.GetHtmlParser()

	selection := query.Find(".s-result-list .result-content .p-name a")
	log.Debug(selection.Size())
	url, _ := selection.Attr("href")
	url = strings.Trim(url, " \t\n")
	sp := spider.NewSpider(NewBookUUDetailProcesser(), "BookUU")
	ip := getProxyIp()
	req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
	if ip != "" {
		req.AddProxyHost(ip)
	}
	pageItems := sp.GetByRequest(req)
	//pageItems := sp.Get("http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn", "html")
	if pageItems == nil || pageItems.GetAll() == nil {
		return
	}
	log.Debug("-----------------------------------spider.Get---------------------------------")
	log.Debug("url\t:\t" + url)
	for name, value := range pageItems.GetAll() {
		p.AddField(name, value)
	}

}

func (s *BookUUListProcesser) Finish() {
	fmt.Printf("TODO:before end spider \r\n")
}

type BookUUDetailProcesser struct {
}

func NewBookUUDetailProcesser() *BookUUDetailProcesser {
	return &BookUUDetailProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (s *BookUUDetailProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		log.Debug(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()
	//获取图书名称 简介
	title := query.Find("#name h2").Text()
	title = strings.Trim(title, " \t\n")

	//获取图书作者，出版社 ，出版时间
	var author, publisher, pubdate string
	detailStr := query.Find(".parameter").Text()
	detailStr = strings.Trim(detailStr, " \t\n")
	//detailStr = strings.Replace(detailStr, " ", "", -1)
	detailStr = strings.Replace(detailStr, "\n", "", -1)
	detailStr = strings.Replace(detailStr, "\t", "", -1)
	//版本
	reg := regexp.MustCompile("印次：\\d+")
	edition := reg.FindString(detailStr)
	edition = strings.Replace(edition, "印次：", "", -1)
	//出版时间
	reg = regexp.MustCompile("出版时间：\\d{4}-\\d{1,2}")
	pubdate = reg.FindString(detailStr)
	pubdate = strings.Replace(pubdate, "出版时间：", "", -1)

	//获取出版社
	reg = regexp.MustCompile("作　者：[\\p{Han}]+")
	author = reg.FindString(detailStr)
	author = strings.Replace(author, "作　者：", "", -1)

	//获取作者
	reg = regexp.MustCompile("出版社：[\\p{Han}]+")
	publisher = reg.FindString(detailStr)
	publisher = strings.Replace(publisher, "出版社：", "", -1)

	//获取isbn
	reg = regexp.MustCompile("(\\d[- ]*){12}[\\d]")
	isbn := reg.FindString(detailStr)
	isbn = strings.Replace(isbn, "-", "", -1)
	isbn = strings.Replace(isbn, " ", "", -1)

	//获取图片url
	url, _ := query.Find(".show-pic img").Attr("src")
	url = strings.Trim(url, " \t\n")

	//获取图书价格
	price := query.Find(".original-price").Text()
	price = strings.Trim(price, " \t\n")
	reg = regexp.MustCompile("\\d+.\\d{2}")
	price = reg.FindString(price)
	if edition != "" {
		edition = "第" + edition + "版"
	}
	p.AddField("title", title)
	p.AddField("remark", "")
	p.AddField("author", author)
	p.AddField("publisher", publisher)
	p.AddField("pubdate", pubdate)
	p.AddField("price", price)
	p.AddField("isbn", isbn)
	p.AddField("image_url", url)
	p.AddField("edition", edition)

}

func (s *BookUUDetailProcesser) Finish() {

}
