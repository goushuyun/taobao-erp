package bookspider

import (
	"regexp"
	"strings"

	log "github.com/goushuyun/log"

	"github.com/hu17889/go_spider/core/common/page"
)

type TaobaoListProcesser struct {
}

func NewTaobaoListProcesser() *TaobaoListProcesser {
	return &TaobaoListProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (s *TaobaoListProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		log.Debug(p.Errormsg())
		return
	}
	query := p.GetHtmlParser()
	content := query.Text()
	reg := regexp.MustCompile("\"category\":\"(\\d*)\",")
	results := reg.FindAllString(content, -1)
	if len(results) <= 0 {
		return
	}
	cids := make(map[string]int)
	for i := 0; i < len(results); i++ {
		category := results[i]
		category = strings.Replace(category, "\"category\":\"", "", -1)
		category = strings.Replace(category, "\",", "", -1)
		cids[category] = cids[category] + 1
	}
	var greaterCid string
	var greaterNum int
	for k, v := range cids {
		if greaterNum == 0 {
			greaterCid = k
			greaterNum = v
			continue
		}
		if v > greaterNum {
			greaterCid = k
			greaterNum = v
		}
	}
	p.AddField("category", greaterCid)
}

func (s *TaobaoListProcesser) Finish() {
}

type TaobaoDetailProcesser struct {
}

func NewTaobaoDetailProcesser() *TaobaoDetailProcesser {
	return &TaobaoDetailProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (s *TaobaoDetailProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		log.Debug(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()
	cid, _ := query.Find("#J_Pinus_Enterprise_Module").Attr("data-catid")
	p.AddField("cid", cid)
}

func (s *TaobaoDetailProcesser) Finish() {

}
