package bookspider

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/wothing/log"

	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
)

type YouLuListProcesser struct {
}

func NewYouLuListProcesser() *YouLuListProcesser {
	return &YouLuListProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (s *YouLuListProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		log.Debug(p.Errormsg())
		return
	}
	query := p.GetHtmlParser()

	selection := query.Find(" .result_list .book_face a")
	url, _ := selection.Attr("href")
	url = strings.Trim(url, " \t\n")
	if url == "" {
		log.Error("errrr")
		return
	}
	url = "http://www.youlu.net" + url
	sp := spider.NewSpider(NewYouLuDetailProcesser(), "YouLu")
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

func (s *YouLuListProcesser) Finish() {
}

type YouLuDetailProcesser struct {
}

func NewYouLuDetailProcesser() *YouLuDetailProcesser {
	return &YouLuDetailProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (s *YouLuDetailProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		log.Debug(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()
	var series_name, page, packing, format, catalog, abstract, author_info, author, publisher, pubdate, title, isbn, price, url, edition string

	//获取书名
	title = query.Find("#name").Text()
	title = replaceSpaceAndLine(title)
	//获取图片
	url, _ = query.Find(".show-pic .pic img").Attr("src")
	isCon := strings.ContainsAny(url, "noPhoto")
	if isCon {
		url = ""
	}
	query.Find("#main-detail-info li").Each(func(i int, s *goquery.Selection) {
		band := s.Text()
		band = replaceSpaceAndLine(band)
		//出版社
		if strings.Contains(band, "出版社：") {
			publisher = strings.Replace(band, "出版社：", "", -1)
		}
		//获取价格
		if strings.Contains(band, "定价：") {
			price = strings.Replace(band, "定价：", "", -1)
			price = strings.Replace(price, "¥", "", -1)
		}
		//获取isbn
		if strings.Contains(band, "ISBN：") {
			isbn = strings.Replace(band, "ISBN：", "", -1)
		}
		//获取出版时间
		if strings.Contains(band, "出版日期：") {
			pubdate = strings.Replace(band, "出版日期：", "", -1)
		}
		//作者
		if strings.Contains(band, "作者：") {
			author = strings.Replace(band, "作者：", "", -1)
		}
		//页数
		if strings.Contains(band, "页数：") {
			page = strings.Replace(band, "页数：", "", -1)
		}

		log.Debug(band)
	})

	//目录
	catalog, _ = query.Find("#J_catalog").Html()
	log.Debug(catalog)
	reg := regexp.MustCompile("<script>[\\w\\W]*</script>")
	a := reg.FindString(catalog)
	catalog = strings.Replace(catalog, a, "", -1)
	catalog = strings.Replace(catalog, "", "", -1)
	catalog = strings.Replace(catalog, "display:none", "", -1)
	//内容简介
	abstract, _ = query.Find("#J_summary").Html()
	//作者简介
	author_info, _ = query.Find("#author_info").Html()
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

func (s *YouLuDetailProcesser) Finish() {

}

func replaceSpaceAndLine(band string) string {
	band = strings.Trim(band, " \t\n")
	band = strings.Replace(band, "\t", "", -1)
	band = strings.Replace(band, "\n", " ", -1)
	band = strings.Replace(band, " ", "", -1)
	return band
}
