package bookspider

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/goushuyun/log"

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
	//处理默认图片
	if url == "http://style.bookuu.com/images/none-picture.png" {
		url = ""
	}
	//获取图书价格
	price := query.Find(".original-price").Text()
	price = strings.Trim(price, " \t\n")
	reg = regexp.MustCompile("\\d+.\\d{2}")
	price = reg.FindString(price)
	if edition != "" {
		edition = "第" + edition + "版"
	}

	var series_name, page, packing, format, catalog, abstract, author_info string
	query.Find(".parameter li").Each(func(i int, s *goquery.Selection) {
		band := s.Text()
		band = strings.Trim(band, " \t\n")
		band = strings.Replace(band, " ", "", -1)
		//丛书名
		if strings.Contains(band, "丛书名：") {
			series_name = strings.Replace(band, "丛书名：", "", -1)
		}
		//页数
		if strings.Contains(band, "页数：") {
			page = strings.Replace(band, "页数：", "", -1)

		}
		//包装
		if strings.Contains(band, "包装：") {
			packing = strings.Replace(band, "包装：", "", -1)
		}
		//开本
		if strings.Contains(band, "开本：") {
			format = strings.Replace(band, "开本：", "", -1)

		}
	})

	//目录
	catalog, _ = query.Find("#J_wrap_5").Html()
	reg = regexp.MustCompile("<a.*</a>")
	a := reg.FindString(catalog)
	catalog = strings.Replace(catalog, a, "", -1)
	catalog = strings.Replace(catalog, "display:none", "", -1)
	//内容简介
	abstract, _ = query.Find("#J_wrap_2").Html()
	//作者简介
	author_info, _ = query.Find("#J_wrap_3").Html()
	log.Debug("==============")
	p.AddField("series_name", series_name)
	p.AddField("catalog", catalog)
	p.AddField("abstract", abstract)
	p.AddField("author_info", author_info)
	p.AddField("page", page)
	p.AddField("packing", packing)
	p.AddField("format", format)
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
