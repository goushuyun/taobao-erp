package bookspider

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/wothing/log"

	"github.com/goushuyun/taobao-erp/misc/key_word_filter"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
)

// "github.com/goushuyun/weixin-golang/misc/bookspider"

type BookSearchResult struct {
	err  error
	book *pb.Book
}

//通过爬虫获取图书信息
func GetBookInfoBySpider(isbn, upload_way string) (book *pb.Book, err error) {
	timeout := make(chan bool)
	result := make(chan BookSearchResult)

	var wg sync.WaitGroup
	go func() {
		time.Sleep(25 * time.Second) // 设置查询超时时间
		timeout <- true
	}()
	go func() {
		wg.Add(1)
		book, err = spiderCoreHandler(isbn, upload_way)
		model := BookSearchResult{book: book, err: err}
		result <- model
	}()
	select {
	case model, ok := <-result:
		if ok {
			book = model.book
			err = model.err
		}
		close(result)
		break
	case <-timeout:
		log.Debug("timeout")
		close(timeout)
		break
	}
	return
}

//构建返回结果
func structData(items *page_items.PageItems, book *pb.Book) {
	//获取数据
	title, _ := items.GetItem("title")
	author, _ := items.GetItem("author")
	publisher, _ := items.GetItem("publisher")
	pubdate, _ := items.GetItem("pubdate")
	price, _ := items.GetItem("price")
	isbn, _ := items.GetItem("isbn")
	edition, _ := items.GetItem("edition")
	image_url, _ := items.GetItem("image_url")
	catalog, _ := items.GetItem("catalog")
	abstract, _ := items.GetItem("abstract")
	series_name, _ := items.GetItem("series_name")
	author_info, _ := items.GetItem("author_info")
	page, _ := items.GetItem("page")
	packing, _ := items.GetItem("packing")
	format, _ := items.GetItem("format")
	var priceFloat float64
	// 处理数据
	if price != "" {
		priceFloat, _ = strconv.ParseFloat(price, 64)
		priceFloat = priceFloat * 100
	}
	// 封装数据
	// 过滤数据
	title = key_word_filter.FilterKeyWords(title)
	title = strings.Replace(title, "'", "\"", -1)
	if title != "" {
		book.Title = title
	}
	if isbn != "" {
		isbn = strings.Replace(isbn, "'", "\"", -1)
		book.Isbn = isbn
	}

	if priceFloat != 0 {
		book.Price = int64(priceFloat)
	}
	if author != "" {
		author = strings.Replace(author, "'", "\"", -1)
		book.Author = author
	}
	if publisher != "" {
		publisher = strings.Replace(publisher, "'", "\"", -1)
		book.Publisher = publisher
	}
	if pubdate != "" {
		book.Pubdate = pubdate
	}
	if image_url != "" {
		book.Image = image_url
	}
	if edition != "" {
		book.Edition = edition
	}
	if catalog != "" {
		catalog = strings.Replace(catalog, "'", "\"", -1)
		book.Catalog = catalog
	}
	if abstract != "" {
		abstract = strings.Replace(abstract, "'", "\"", -1)
		book.Abstract = abstract

	}
	if series_name != "" {
		book.SeriesName = series_name
	}
	if author_info != "" {
		author_info = strings.Replace(author_info, "'", "\"", -1)
		book.AuthorIntro = author_info
	}
	if page != "" {
		book.Page = page
	}
	if packing != "" {
		book.Packing = packing
	}
	if format != "" {
		book.Format = format
	}

	return
}

func spiderCoreHandler(isbn, upload_way string) (book *pb.Book, err error) {
	book = &pb.Book{}
	isbn = strings.Replace(isbn, "-", "", -1)
	isbn = strings.Replace(isbn, " ", "", -1)
	var num int32
	if upload_way == "batch" {
		num = rand.Int31n(3)
	} else {
		num = rand.Int31n(1)
	}
	log.Debugf("===上传类型：%s========停留秒数：%d", upload_way, num)
	time.Sleep(time.Duration(num) * time.Second)
	ip := getProxyIp()

	//首先从当当上获取图书信息
	book.SourceInfo = "dangdang"
	sp := spider.NewSpider(NewDangDangListProcesser(), "spiderDangDangList")
	baseURL := "http://search.dangdang.com/?key=ISBN&ddsale=1"
	url := strings.Replace(baseURL, "ISBN", isbn, -1)
	req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)

	log.Debug(url)
	if ip != "" {
		req.AddProxyHost(ip)
	}

	pageItems := sp.GetByRequest(req)
	//没爬到数据
	if pageItems == nil || len(pageItems.GetAll()) <= 0 {
		log.Debug("dangdang no data")
	} else {
		structData(pageItems, book)
		if book.Isbn != "" && isbn == book.Isbn && book.Price != 0 && book.Title != "" && book.Publisher != "" {
			//如果获取到数据，返回
			log.Debugf("%+v", book)
			return
		}

	}

	//从京东上获取图书信息
	book.SourceInfo = "jd"
	sp = spider.NewSpider(NewJDListProcesser(), "spiderJDList")
	baseURL = "https://search.jd.com/Search?keyword=ISBN&enc=utf-8&qrst=1&rt=1&stop=1&vt=2&wq=ISBN&psort=1&wtype=1&click=1"
	url = strings.Replace(baseURL, "ISBN", isbn, -1)
	req = request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)

	if ip != "" {
		req.AddProxyHost(ip)
	}
	pageItems = sp.GetByRequest(req)
	//没爬到数据
	if pageItems == nil || len(pageItems.GetAll()) <= 0 {
		log.Debug("jd no data")
	} else {
		structData(pageItems, book)
		if book.Isbn != "" && isbn == book.Isbn && book.Price != 0 && book.Title != "" {
			//如果获取到数据，返回
			log.Debugf("%+v", book)
			return
		}

	}

	book.SourceInfo = "youlu"
	sp = spider.NewSpider(NewYouLuListProcesser(), "youlu")
	baseURL = "http://www.youlu.net/search/result3/?isbn=ISBN"
	url = strings.Replace(baseURL, "ISBN", isbn, -1)
	req = request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
	if ip != "" {
		req.AddProxyHost(ip)
	}
	pageItems = sp.GetByRequest(req)
	//没爬到数据
	if pageItems == nil || len(pageItems.GetAll()) <= 0 {
		log.Debug("youlu no data")
	} else {
		structData(pageItems, book)
		if book.Isbn != "" && isbn == book.Isbn && book.Price != 0 && book.Title != "" {
			//如果获取到数据，返回
			log.Debugf("%+v", book)
			return
		}

	}
	return nil, nil
}

//通过爬虫获取图书信息
func GetBookTaobaoCategory(isbn string) (category string) {
	timeout := make(chan bool)
	result := make(chan string)

	var wg sync.WaitGroup
	go func() {
		time.Sleep(10 * time.Second) // 设置查询超时时间
		timeout <- true
	}()
	go func() {
		wg.Add(1)
		sp := spider.NewSpider(NewTaobaoListProcesser(), "taobao")
		baseUrl := "https://s.taobao.com/search?q=ISBN&imgfile=&commend=all&ssid=s5-e&search_type=item&sourceId=tb.index&spm=a21bo.50862.201856-taobao-item.1&ie=utf8&initiative_id=tbindexz_20170818"
		url := strings.Replace(baseUrl, "ISBN", isbn, -1)
		req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
		pageItems := sp.GetByRequest(req)
		if pageItems == nil || len(pageItems.GetAll()) <= 0 {
			result <- ""
		} else {
			cid, _ := pageItems.GetItem("category")
			result <- cid
		}

	}()
	select {
	case model, ok := <-result:
		if ok {
			category = model
		}
		close(result)
		break
	case <-timeout:
		log.Debug("timeout")
		close(timeout)
		break
	}
	return
}
