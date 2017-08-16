//
package bookspider

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/goushuyun/taobao-erp/misc"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
	log "github.com/wothing/log"
	"gopkg.in/iconv.v1"
)

const ProxyServer = "proxy.abuyun.com:9020"

type ProxyAuth struct {
	License   string
	SecretKey string
}

func (p ProxyAuth) ProxyClient() http.Client {
	proxyURL, _ := url.Parse("http://" + p.License + ":" + p.SecretKey + "@" + ProxyServer)
	return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
}

func TestSpiderDangdangList(t *testing.T) {
	isbn := "9787301063453"
	sp := spider.NewSpider(NewDangDangListProcesser(), "spiderDangDangList")
	baseURL := "http://search.dangdang.com/?key=ISBN&ddsale=1"
	url := strings.Replace(baseURL, "ISBN", isbn, -1)
	req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
	pageItems := sp.GetByRequest(req)

	//没爬到数据
	if pageItems == nil || len(pageItems.GetAll()) <= 0 {
		log.Debug("no matches found!")
		return
	}
	for name, value := range pageItems.GetAll() {
		log.Debug(name + "\t:\t" + value)
	}
}

func TestSpiderJDList(t *testing.T) {
	isbn := "9787535492838"
	sp := spider.NewSpider(NewJDListProcesser(), "spiderJDList")
	baseURL := "https://search.jd.com/Search?keyword=ISBN&enc=utf-8&qrst=1&rt=1&stop=1&vt=2&wq=ISBN&psort=1&wtype=1&click=1"
	url := strings.Replace(baseURL, "ISBN", isbn, -1)
	req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
	pageItems := sp.GetByRequest(req)
	//pageItems := sp.Get("http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn", "html")

	//没爬到数据
	if pageItems == nil || len(pageItems.GetAll()) <= 0 {
		log.Debug("no matches found!")
		return
	}
	for name, value := range pageItems.GetAll() {
		log.Debug(name + "\t:\t" + value)
	}
}
func TestSpiderDangdangDetail(t *testing.T) {
	sp := spider.NewSpider(NewDangDangDetailProcesser(), "spiderDangDangDetail")
	req := request.NewRequest("http://product.dangdang.com/24170700.html", "html", "", "GET", "", nil, nil, nil, nil)

	pageItems := sp.GetByRequest(req)

	url := pageItems.GetRequest().GetUrl()
	log.Debug("-----------------------------------spider.Get---------------------------------")
	log.Debug("url\t:\t" + url)
	for name, value := range pageItems.GetAll() {
		log.Debug(name + "\t:\t" + value)
	}
}

func TestSpiderBookUUList(t *testing.T) {
	isbn := "9787559402585"
	sp := spider.NewSpider(NewBookUUListProcesser(), "BookUUlist")
	baseUrl := "http://search.bookuu.com/AdvanceSearch.php?isbn=ISBN&sm=&zz=&cbs=&dj_s=&dj_e=&bkj_s=&bkj_e=&layer2=&zk=0&cbrq_n=2017&cbrq_y=&cbrq_n1=2017&cbrq_y1=&sjsj=0&orderby=&layer1=1"
	url := strings.Replace(baseUrl, "ISBN", isbn, -1)
	req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)

	pageItems := sp.GetByRequest(req)
	for name, value := range pageItems.GetAll() {
		log.Debug(name + "\t:\t" + value)
	}
}

func TestSpiderCaiCoolList(t *testing.T) {
	isbn := "9787562165576"
	sp := spider.NewSpider(NewCaiCoolListProcesser(), "CaiCoolList")
	baseUrl := "http://www.caicool.cn/search?keywords=ISBN&typesMark=0&typesCode=-1&switchMark=0"
	url := strings.Replace(baseUrl, "ISBN", isbn, -1)
	req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)

	pageItems := sp.GetByRequest(req)
	for name, value := range pageItems.GetAll() {
		log.Debug(name + "\t:\t" + value)
	}
}

func TestSpiderYouluList(t *testing.T) {
	isbn := "9787540455958"
	sp := spider.NewSpider(NewYouLuListProcesser(), "youlu")
	baseUrl := "http://www.youlu.net/search/result3/?isbn=ISBN&publisherName=&author=&bookName="
	url := strings.Replace(baseUrl, "ISBN", isbn, -1)
	req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)

	pageItems := sp.GetByRequest(req)
	for name, value := range pageItems.GetAll() {
		log.Debug(name + "\t:\t" + value)
	}
}
func TestGetBookInfo(t *testing.T) {
	book, _ := GetBookInfoBySpider("9787121266942", "")
	println("-----------------------------------OOOOOOM---------------------------------")
	log.Debugf("%#v", book)
	log.Debug("-----------------------------------OOOOOOM---------------------------------")

}
func TestRegular(t *testing.T) {
	detailStr := "https://item.jd.com/11020022.html"
	reg := regexp.MustCompile("/\\d*\\.")
	log.Debug(reg.FindString(detailStr))

}
func TestGbkToUtf(t *testing.T) {
	log.Debug(111)
	res, err := http.Get("https://dx.3.cn/desc/12155418?cdn=2&callback=showdesc")
	if err != nil {
		log.Debug(1)
		log.Debug(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

		log.Debug(err)
	}
	convStr := strings.Replace(string(body), "\\\"", "'", -1)
	reader := bytes.NewReader([]byte(convStr))
	cd, err := iconv.Open("utf-8", "gbk") // convert gbk to utf8
	if err != nil {
		fmt.Println("iconv.Open failed!")
	}
	defer cd.Close()

	utfBody := iconv.NewReader(cd, reader, 0)
	if err != nil {
		log.Debug(err)

	}
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Debug(err)
	}
	log.Debug(111)

	catalog, _ := doc.Find("#detail-tag-id-6").Html()
	//catalog, _ := doc.Html()

	log.Debug(catalog)
}

func TestUnicodeToUtf(t *testing.T) {
	log.Debug(111)
	res, err := http.Get("http://product.dangdang.com/index.php?r=callback%2Fdetail&productId=1900465424&templateType=publish&describeMap=&shopId=0&categoryPath=01.49.01.18.00.00")
	if err != nil {
		log.Debug(1)
		log.Debug(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debug(2)
		log.Debug(err)
	}
	context := misc.UnicodeToUtf8(string(body))

	reader := bytes.NewReader([]byte(context))

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Debug(3)
		log.Debug(err)
	}
	catalog, err := doc.Find("#catalog-textarea").Html()
	if err != nil {
		log.Debug(3)
		log.Debug(err)
	}

	content, err := doc.Find("#content").Html()
	if err != nil {
		log.Debug(3)
		log.Debug(err)
	}
	log.Debug("=================")
	log.Debug(misc.UnicodeToUtf8(catalog))
	log.Debug(misc.UnicodeToUtf8(content))

}

func TestAbuyun(t *testing.T) {
	targetURI := "http://ip.chinaz.com/"
	//targetURI := "http://www.abuyun.com/switch-ip"
	//targetURI := "http://www.abuyun.com/current-ip"
	reg := regexp.MustCompile("((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)")

	// 初始化 proxy http client
	client := ProxyAuth{License: "H2YYNX817619N32D", SecretKey: "73FAB0143E36EF3D"}.ProxyClient()

	request, _ := http.NewRequest("GET", targetURI, bytes.NewBuffer([]byte(``)))

	// 切换IP (只支持 HTTP)
	request.Header.Set("Proxy-Switch-Ip", "yes")

	response, err := client.Do(request)

	if err != nil {
		panic("failed to connect: " + err.Error())
	} else {
		bodyByte, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("读取 Body 时出错", err)
			return
		}
		response.Body.Close()

		body := string(bodyByte)

		fmt.Println("Response Status:", response.Status)
		fmt.Println("Response Header:", response.Header)
		fmt.Println("Response Body:\n", body)

		ip := reg.FindString(string(body))
		fmt.Printf("\n代理ip：%s\n", ip) //打印
	}

}

func TestJdAnaly(t *testing.T) {
	priceUrl := "https://upload.taobao.com/auction/json/get_category.htm?isB2CSeller=false&is3CSeller=false&isSynergicSeller=false&isAliMall=false&isFinancingSeller=false&isAssets=false&isTaoDianDian=false&fenxiaoProduct=&_input_charset=utf-8&keyword=9787115399922"
	// reg := regexp.MustCompile("/\\d*\\.")
	// productId := reg.FindString(productUrl)
	// productId = strings.Replace(productId, ".", "", -1)
	// productId = strings.Replace(productId, "/", "", -1)
	// log.Debug("productId========", productId)
	// priceUrl = strings.Replace(priceUrl, "PRODUCTID", productId, -1)
	log.Debug("priceUrl========", priceUrl)
	resp, err := http.Post(priceUrl,
		"application/text/html",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	log.Debug(string(body))

}
