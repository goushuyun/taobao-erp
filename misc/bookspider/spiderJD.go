package bookspider

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
	log "github.com/wothing/log"
)

type JDListProcesser struct {
}

func NewJDListProcesser() *JDListProcesser {
	return &JDListProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
func (s *JDListProcesser) Process(p *page.Page) {
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

	selection := query.Find("#J_goodsList ul li .p-img")
	if selection.Size() > 0 {

	}

	findUrl, _ := selection.Find("a").Attr("href")
	log.Debug("-----------------------------------spider.Url---------------------------------")
	log.Debug(findUrl)
	log.Debug("-----------------------------------spider.Url---------------------------------")
	findUrl = strings.Trim(findUrl, " \t\n")
	if findUrl == "" {
		log.Error("京东无数据")
		return
	}
	findUrl = "https:" + findUrl
	sp := spider.NewSpider(NewJDDetailProcesser(), "JDDetail")
	ip := getProxyIp()
	var req *request.Request
	req = request.NewRequest(findUrl, "html", "", "GET", "", nil, nil, nil, nil)
	if ip != "" {
		req.AddProxyHost(ip)
	}
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

func (s *JDListProcesser) Finish() {

}

type JDDetailProcesser struct {
}

func NewJDDetailProcesser() *JDDetailProcesser {
	return &JDDetailProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (s *JDDetailProcesser) Process(p *page.Page) {
	var author, publisher, pubdate, price string
	if !p.IsSucc() {
		log.Debug(p.Errormsg())
		return
	}

	productUrl := p.GetRequest().GetUrl()
	log.Debug("URL========", productUrl)
	priceUrl := "http://p.3.cn/prices/mgets?skuIds=J_PRODUCTID"
	reg := regexp.MustCompile("/\\d*\\.")
	productId := reg.FindString(productUrl)
	productId = strings.Replace(productId, ".", "", -1)
	productId = strings.Replace(productId, "/", "", -1)

	log.Debug("productId========", productId)
	if productId == "" {
		log.Debug("京东无数据")
		return
	}
	priceUrl = strings.Replace(priceUrl, "PRODUCTID", productId, -1)
	log.Debug("priceUrl========", priceUrl)
	ipStr := getProxyIp()
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(ipStr) //根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	resp, err := client.Get(priceUrl) //请求并获取到对象,使用代理
	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //取出主体的内容
	if err != nil {
		log.Error(err)
		return
	}

	log.Debug(string(body))
	//获取价格
	var param []map[string]string
	err = json.Unmarshal(body, &param)
	if err != nil {
		log.Debug(err)
		return
	} else {
		price = param[0]["m"]
		if price == "" {
			return
		}
	}
	query := p.GetHtmlParser()
	//获取图书名称 简介
	title := query.Find("#name h1").Text()
	title = strings.Trim(title, " \t\n")

	remark := query.Find("#p-ad").Text()
	remark = strings.Trim(remark, " \t\n")

	//获取图书作者，出版社 ，出版时间
	detailStr := query.Find("#parameter2").Text()
	detailStr = strings.Trim(detailStr, " \t\n")
	detailStr = strings.Replace(detailStr, "\n", "", -1)
	detailStr = strings.Replace(detailStr, " ", "", -1)
	log.Debug(detailStr)

	author = query.Find("#p-author").Text()
	author = strings.Trim(author, " \t\n")

	reg = regexp.MustCompile("出版社：.*ISBN")
	publisher = reg.FindString(detailStr)
	publisher = strings.Replace(publisher, "出版社：", "", -1)
	publisher = strings.Replace(publisher, "ISBN", "", -1)

	reg = regexp.MustCompile("\\d{4}[\\p{Han}-]{1}\\d{2}[\\p{Han}]{0,1}")
	pubdate = reg.FindString(detailStr)

	//获取isbn
	reg = regexp.MustCompile("(\\d[- ]*){12}[\\d]")
	isbn := reg.FindString(detailStr)
	isbn = strings.Replace(isbn, "-", "", -1)
	isbn = strings.Replace(isbn, " ", "", -1)

	reg = regexp.MustCompile("版 次：[\\d]+")
	edition := reg.FindString(detailStr)
	edition = strings.Replace(edition, "版 次：", "", -1)
	if edition != "" {
		edition = "第" + edition + "版"
	}
	//获取图片url
	url, _ := query.Find("#spec-n1 img").Attr("src")
	url = strings.Trim(url, " \t\n")
	if url != "" {
		url = "https:" + url
	}
	p.AddField("title", title)
	p.AddField("remark", remark)
	p.AddField("author", author)
	p.AddField("publisher", publisher)
	p.AddField("pubdate", pubdate)
	p.AddField("price", price)
	p.AddField("isbn", isbn)
	p.AddField("image_url", url)
	p.AddField("edition", edition)

}

func (s *JDDetailProcesser) Finish() {

}
