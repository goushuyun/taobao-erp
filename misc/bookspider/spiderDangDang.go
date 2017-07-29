package bookspider

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
	log "github.com/wothing/log"
)

type DangDangListProcesser struct {
}

func NewDangDangListProcesser() *DangDangListProcesser {
	return &DangDangListProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
func (s *DangDangListProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		log.Debug(p.Errormsg())
		return
	}
	//p.AddTargetRequestsWithProxy(p., respType, proxyHost)
	header := http.Header{}
	header.Add("User-Agent", "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0)")
	header.Add("Content-Type", "application/json")
	header.Add("Authentization", "sdfsdfdsfsdsdfsdfsdfsdfsd")
	p.SetHeader(header)
	query := p.GetHtmlParser()

	selection := query.Find(".bigimg li")
	if selection.Size() > 0 {

	}

	findUrl, _ := selection.Find("a.pic").Attr("href")
	findUrl = strings.Trim(findUrl, " \t\n")
	sp := spider.NewSpider(NewDangDangDetailProcesser(), "DangDangDetail")
	req := request.NewRequest(findUrl, "html", "", "GET", "", nil, nil, nil, nil)

	pageItems := sp.GetByRequest(req)
	if pageItems == nil || pageItems.GetAll() == nil {
		return
	}
	log.Debug("-----------------------------------spider.Get---------------------------------")
	log.Debug("url\t:\t" + findUrl)
	for name, value := range pageItems.GetAll() {
		p.AddField(name, value)
	}
}

func (s *DangDangListProcesser) Finish() {

}

type DangDangDetailProcesser struct {
}

func NewDangDangDetailProcesser() *DangDangDetailProcesser {
	return &DangDangDetailProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (s *DangDangDetailProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		log.Debug(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()
	//获取图书名称 简介
	title := query.Find(".sale_box_left .name_info h1").Text()
	title = strings.Trim(title, " \t\n")
	//获取图书作者，出版社 ，出版时间
	var author, publisher, pubdate string
	author = query.Find(".messbox_info span a[dd_name='作者']").Text()
	publisher = query.Find(".messbox_info span a[dd_name='出版社']").Text()
	pubdate = query.Find(".messbox_info span:nth-child(3)").Text()

	author = strings.Trim(author, " \t\n")
	publisher = strings.Trim(publisher, " \t\n")
	pubdate = strings.Trim(pubdate, " \t\n")
	reg := regexp.MustCompile("\\d{4}[\\p{Han}]{1}\\d{2}[\\p{Han}]{1}")
	pubdate = reg.FindString(pubdate)
	//获取图书价格
	price := query.Find("#original-price").Text()
	price = strings.Trim(price, " \t\n")
	reg = regexp.MustCompile("(\\d+).\\d{2}")
	price = reg.FindString(price)
	//获取isbn
	isbnStr := query.Find("#detail_describe .key li").Text()
	reg = regexp.MustCompile("(\\d[- ]*){12}[\\d]")
	isbn := reg.FindString(isbnStr)
	isbn = strings.Replace(isbn, "-", "", -1)
	isbn = strings.Replace(isbn, " ", "", -1)
	reg = regexp.MustCompile("版 次：[\\d]+")
	edition := reg.FindString(isbnStr)
	edition = strings.Replace(edition, "版 次：", "", -1)
	//获取图片url
	url, _ := query.Find("#largePicDiv #largePic").Attr("src")
	url = strings.Trim(url, " \t\n")
	if edition != "" {
		edition = "第" + edition + "版"
	}

	p.AddField("title", title)
	p.AddField("author", author)
	p.AddField("publisher", publisher)
	p.AddField("edition", edition)
	log.Debug("==============")
	content := query.Find("#detail_describe .key").Text()
	log.Debug(content)
	log.Debug("==============")

	//丛书名

	//目录
	//内容简介
	//页数
	//包装
	//开本
	//作者简介
	p.AddField("pubdate", pubdate)
	p.AddField("price", price)
	p.AddField("isbn", isbn)
	p.AddField("image_url", url)

}

func (s *DangDangDetailProcesser) Finish() {

}
