package misc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestRegexp(t *testing.T) {
	matchStr := "11.168.2.2"
	isMatch, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", matchStr)
	fmt.Printf(matchStr+" is ip ?  answer is %b", isMatch)
}

func TestSpider(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http get error")
	}
	src := string(body)

	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	fmt.Println(strings.TrimSpace(src))
}
func TestCmd(t *testing.T) {
	HOME := os.Getenv("POSTGRES_PORT_5432_TCP_ADDR")
	fmt.Println(HOME)
}
