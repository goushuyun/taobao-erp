//
package bookspider

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/goushuyun/log"
)

func getOrderNo() string {
	return "d64615fa08c3dfea28fa9c0a1fbc3791&random=true&sep=3"
}
func getProxyIp() string {
	orderNo := getOrderNo()
	url := "http://api.ip.data5u.com/dynamic/get.html?order=" + orderNo
	resp, err := http.Post(url,
		"application/text/html",
		strings.NewReader("name=cjb"))
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Error(err)
		return ""
	}
	ipStr := string(body)
	reg := regexp.MustCompile("((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)")
	ip := reg.FindString(string(body))

	if ip == "" {
		ipStr = ip
	}
	ipStr = strings.TrimSpace(ipStr)
	log.Debug(ipStr)
	return "http://" + ipStr
}
