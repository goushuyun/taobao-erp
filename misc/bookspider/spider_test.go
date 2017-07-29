//
package bookspider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"

	log "github.com/wothing/log"

	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
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

func TestSpiderJDList(t *testing.T) {
	isbn := "9787301091319"
	sp := spider.NewSpider(NewJDListProcesser(), "spiderJDList")
	baseURL := "https://search.jd.com/Search?keyword=ISBN&enc=utf-8&wq=ISBN&pvid=3d3aefa8a0904ef1b08547fb69f57ae7"
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
	isbn := "9787559602404"
	sp := spider.NewSpider(NewBookUUListProcesser(), "BookUUlist")
	baseUrl := "http://search.bookuu.com/AdvanceSearch.php?isbn=ISBN&sm=&zz=&cbs=&dj_s=&dj_e=&bkj_s=&bkj_e=&layer2=&zk=0&cbrq_n=2017&cbrq_y=&cbrq_n1=2017&cbrq_y1=&sjsj=0&orderby=&layer1=1"
	url := strings.Replace(baseUrl, "ISBN", isbn, -1)
	req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)

	pageItems := sp.GetByRequest(req)
	for name, value := range pageItems.GetAll() {
		log.Debug(name + "\t:\t" + value)
	}
}

// func TestGetBookInfo(t *testing.T) {
// 	book, _ := GetBookInfoBySpider("9787508622019", "")
// 	println("-----------------------------------OOOOOOM---------------------------------")
// 	fmt.Printf("%#v", book)
// 	log.Debug("-----------------------------------OOOOOOM---------------------------------")
//
// }
func TestRegular(t *testing.T) {
	detailStr := "https://item.jd.com/11020022.html"
	reg := regexp.MustCompile("/\\d*\\.")
	log.Debug(reg.FindString(detailStr))

}
func TestProxyIp(t *testing.T) {
	orderNo := getOrderNo()
	url := "http://api.ip.data5u.com/dynamic/get.html?order=" + orderNo
	resp, err := http.Post(url,
		"application/text/html",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Error(err)
		return
	}

	reg := regexp.MustCompile("((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)")
	ip := reg.FindString(string(body))
	log.Debug(string(body))
	log.Debug(ip)

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
	priceUrl := "http://p.3.cn/prices/mgets?skuIds=J_12460649031"
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
	var price string
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	//获取价格
	var param []map[string]string
	log.Debug(string(body))
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

	log.Debug("==============:%s", price)
}
