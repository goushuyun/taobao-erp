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

type CaiCoolListProcesser struct {
}

func NewCaiCoolListProcesser() *CaiCoolListProcesser {
	return &CaiCoolListProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (s *CaiCoolListProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		log.Debug(p.Errormsg())
		return
	}
	query := p.GetHtmlParser()

	selection := query.Find(".cpyp-lb li .sslb-smz a")
	url, _ := selection.Attr("href")
	url = strings.Trim(url, " \t\n")
	if url == "" {
		return
	}
	sp := spider.NewSpider(NewCaiCoolDetailProcesser(), "CaiCool")
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
	for name, value := range pageItems.GetAll() {
		p.AddField(name, value)
	}

}

func (s *CaiCoolListProcesser) Finish() {
}

type CaiCoolDetailProcesser struct {
}

func NewCaiCoolDetailProcesser() *CaiCoolDetailProcesser {
	return &CaiCoolDetailProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (s *CaiCoolDetailProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		log.Debug(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()
	var series_name, page, packing, format, catalog, abstract, author_info, author, publisher, pubdate, title, isbn, price, url, edition string
	//获取书名

	//获取出版社
	//获取作者
	//获取isbn
	//获取出版时间
	//获取价格
	//获取图片
	query.Find(".step-1 .lbwkz li").Each(func(i int, s *goquery.Selection) {
		band := s.Text()
		band = strings.Trim(band, " \t\n")
		band = strings.Replace(band, "\t", "", -1)
		band = strings.Replace(band, "\n", "", -1)
		band = strings.Replace(band, " ", "", -1)
		//丛书名
		if strings.Contains(band, "丛书名：") {
			series_name = strings.Replace(band, "丛书名：", "", -1)
		}
		//页数
		if strings.Contains(band, "页码：") {
			page = strings.Replace(band, "页码：", "", -1)

		}
		//包装
		if strings.Contains(band, "包装：") {
			packing = strings.Replace(band, "包装：", "", -1)
		}
		//开本
		if strings.Contains(band, "开本：") {
			format = strings.Replace(band, "开本：", "", -1)

		}

		log.Debug(band)
	})

	//目录
	catalog, _ = query.Find(".step-5").Html()
	reg := regexp.MustCompile("<script>[\\w\\W]*</script>")
	a := reg.FindString(catalog)
	catalog = strings.Replace(catalog, a, "", -1)
	catalog = strings.Replace(catalog, "", "", -1)
	catalog = strings.Replace(catalog, "display:none", "", -1)
	//内容简介
	abstract, _ = query.Find(".step-3").Html()
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

func (s *CaiCoolDetailProcesser) Finish() {

}
