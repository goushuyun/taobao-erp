package bookspider

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	iconv "gopkg.in/iconv.v1"

	"github.com/PuerkitoBio/goquery"
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
	var author, publisher, pubdate, price, series_name, page, packing, format, catalog, abstract, author_info, edition, isbn string
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
	//获取商品详情url
	detailUrl := strings.Replace("https://dx.3.cn/desc/PRODUCTID?cdn=2&callback=showdesc", "PRODUCTID", productId, -1)
	//获取商品价格url
	priceUrl = strings.Replace(priceUrl, "PRODUCTID", productId, -1)

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

	author = query.Find("#p-author").Text()
	author = strings.Trim(author, " \t\n")

	//获取图片url
	url, _ := query.Find("#spec-n1 img").Attr("src")
	url = strings.Trim(url, " \t\n")
	if url != "" {
		url = "https:" + url
	}

	query.Find("#parameter2 li").Each(func(i int, s *goquery.Selection) {
		band := s.Text()
		band = strings.Trim(band, "\t\n")
		band = strings.Replace(band, "\n", "", -1)
		band = strings.Replace(band, "\t", "", -1)
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
		//版 次
		if strings.Contains(band, "版次：") {
			edition = strings.Replace(band, "版次：", "", -1)
			edition = "第" + edition + "版"
		}
		//出版社
		if strings.Contains(band, "出版社：") {
			publisher = strings.Replace(band, "出版社：", "", -1)
		}
		//出版日期
		if strings.Contains(band, "出版时间：") {
			pubdate = strings.Replace(band, "出版时间：", "", -1)
		}
		//isbn
		if strings.Contains(band, "ISBN：") {
			isbn = strings.Replace(band, "ISBN：", "", -1)
		}

	})
	res, err := client.Get(detailUrl)
	if err != nil {
		log.Debug(err)
		return
	}

	//把gbk编码转换成utf-8编码
	cd, err := iconv.Open("utf-8", "gbk") // convert gbk to utf8
	if err != nil {
		log.Error(err)
	} else {

		defer cd.Close()
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {

			log.Error(err)
		} else {
			//构建goquery
			convStr := strings.Replace(string(body), "\\\"", "'", -1)
			reader := bytes.NewReader([]byte(convStr))
			cd, err := iconv.Open("utf-8", "gbk") // convert gbk to utf8
			if err != nil {
				log.Error("iconv.Open failed!")
			}
			defer cd.Close()

			utfBody := iconv.NewReader(cd, reader, 0)
			if err != nil {
				log.Debug(err)

			}
			doc, err := goquery.NewDocumentFromReader(utfBody)
			if err != nil {
				log.Error(err)
			} else {
				//目录
				catalog, _ = doc.Find("#detail-tag-id-6").Html()
				reg = regexp.MustCompile("<a.*</a>")
				a := reg.FindString(catalog)
				catalog = strings.Replace(catalog, a, "", -1)
				catalog = strings.Replace(catalog, "display:none", "", -1)
				catalog = strings.Replace(catalog, "\\n", "", -1)
				//内容简介
				abstract, _ = doc.Find("#detail-tag-id-3").Html()
				abstract = strings.Replace(abstract, "\\n", "", -1)

				//作者简介
				author_info, _ = doc.Find("#detail-tag-id-4").Html()
				author_info = strings.Replace(author_info, "\\n", "", -1)

			}
		}

	}

	p.AddField("series_name", series_name)
	p.AddField("page", page)
	p.AddField("packing", packing)
	p.AddField("format", format)
	p.AddField("catalog", catalog)
	p.AddField("abstract", abstract)
	p.AddField("author_info", author_info)
	p.AddField("title", title)
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
