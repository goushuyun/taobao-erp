package service

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"strings"

	"github.com/goushuyun/weixin-golang/pb"

	"github.com/tealeg/xlsx"
	"github.com/wothing/log"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var cli *kodo.Client
var uploader kodocli.Uploader

const (
	ak = "PRBrdNUf1m07enHI26CgpG7z3O8YLYXgoCE-1P2S"
	sk = "PtCRtf7xyY0P6mBaQARL6KeYW6dVdglFzIz3BiQS"
)

type qiniuZone struct {
	public bool
	bucket string
	domain string
	imgsuf string
	thmsuf string
}

var zones = map[string]*qiniuZone{
	"Test": &qiniuZone{
		public: true,
		bucket: "bookcloud-test",
		domain: "http://images.goushuyun.cn/",
		imgsuf: "",
		thmsuf: "-th",
	},
	"Public": &qiniuZone{
		public: true,
		bucket: "bookcloud-public",
		domain: "http://onv97cvt2.bkt.clouddn.com/",
		imgsuf: "",
		thmsuf: "-th",
	},
}

func init() {
	kodo.SetMac(ak, sk)
	zone := 0
	cli = kodo.New(zone, nil)
	uploader = kodocli.NewUploader(zone, nil) // 构建一个uploader
}

// put Excel file
func PutExcelFile(excel *xlsx.File) (key string, err error) {
	f, err := ioutil.TempFile("", "tmp")
	if err != nil {
		log.Error(err)
		return "", err
	}
	defer os.Remove(f.Name())

	err = excel.Write(f)
	if err != nil {
		log.Error(err)
		return "", err
	}

	file_name := f.Name()[strings.LastIndex(f.Name(), "/")+1:] + ".xlsx"

	key = "tmp_30days/" + file_name
	token, _ := makeToken(pb.MediaZone_Test, key)

	fi, err := f.Stat()
	if err != nil {
		log.Error(err)
		return "", err
	}

	err = uploader.Rput(context.Background(), nil, token, key, f, fi.Size(), nil)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return key, nil
}

func upload(url string, zone pb.MediaZone, appid string) (key string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return "", err
	}

	mediaType, mediaParams, _ := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	realName := mediaParams["filename"]
	log.Debugf("MediaType=%s, filename=%s, size=%d", mediaType, realName, resp.ContentLength)

	key = appid + "/" + realName

	token, url := makeToken(zone, key)

	log.Debugf("Token : %s, url : %s", token, url)

	if resp.ContentLength <= 1<<22 {
		log.Debug("<<little>> \n")

		putLittleFile(token, key, resp.Body, resp.ContentLength)
	} else {
		log.Debug("<<large>> \n")
		putLargeFile(token, key, resp.Body, resp.ContentLength)
	}

	return key, nil
}

func uploadLocal(filepath string, zone pb.MediaZone, filename string) (err error) {
	token, url := makeToken(zone, filename)

	err = uploader.PutFile(nil, nil, token, filename, filepath, nil)
	//打印出错信息
	if err != nil {
		log.Error(err)
		return err
	}

	log.Debug(url)
	return nil
}

func putLittleFile(token, filename string, body io.ReadCloser, length int64) {

	err := uploader.Put(nil, nil, token, filename, body, length, nil)
	if err != nil {
		log.Error(err)
		return
	}
	body.Close()
	log.Infof("file %s upload complete", filename)
}

func tryPut(token, key string, file *os.File, filesize int64) {
	err := uploader.Rput(context.Background(), nil, token, key, file, filesize, nil)
	if err != nil {
		log.Error(err)
		return
	}
}

func putLargeFile(token, filename string, body io.ReadCloser, length int64) {
	temp, err := ioutil.TempFile("", filename[strings.Index(filename, "/")+1:])
	if err != nil {
		log.Error(err)
		return
	}
	written, err := io.Copy(temp, body)
	if err != nil {
		log.Error(err)
		return
	}
	body.Close()

	log.Debugf("size=%d, temp=%s, written=%d, %p", length, temp.Name(), written, &uploader)

	err = uploader.Rput(context.Background(), nil, token, filename, temp, written, nil)
	if err != nil {
		log.Error(err)
		return
	}
	temp.Close()
	log.Infof("file %s upload complete", filename)

	os.Remove(temp.Name())
}

func makeToken(zone pb.MediaZone, filename string) (token, url string) {
	var c *qiniuZone
	if zone >= pb.MediaZone_Test && zone <= pb.MediaZone_Public {
		c = zones[zone.String()]
	} else {
		c = zones["Test"]
	}

	scope := c.bucket

	if filename != "" {
		scope = scope + ":" + filename
	}

	log.Debugf("The scope is : ", scope)

	policy := &kodo.PutPolicy{
		Scope:   scope,
		Expires: 600,
	}

	token = cli.MakeUptoken(policy)
	url = c.domain + filename

	return
}

func FetchImg(zone pb.MediaZone, url, key string) (string, error) {
	z := zones[zone.String()]

	log.Debugf("The zone is : %+v", z)

	p := cli.Bucket(z.bucket)
	err := p.Fetch(nil, key, url)
	if err != nil {
		return "", err
	}
	log.Debugf("Fetching successful, The url is: ", z.domain+key)
	return z.domain + key, nil
}

// refresh URL cache
type RefreshReq struct {
	Urls []string `json:"urls"`
}

type RefreshResp struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func RefreshURLCache(urls []string) error {
	if len(urls) > 100 {
		return errors.New("The amount of url can't over 100")
	}

	// the url to refresh cache
	post_url := "http://fusion.qiniuapi.com/v2/tune/refresh"

	// urls array to json
	urls_bytes, err := json.Marshal(RefreshReq{Urls: urls})
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debugf("The urls JSON is %s", urls_bytes)

	// send post request
	req, err := http.NewRequest("POST", post_url, bytes.NewBuffer(urls_bytes))
	req.Header.Set("Authorization", "QBox "+GenAccessToken("/v2/tune/refresh\n"))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Debugf("%s", body)
	respStruct := &RefreshResp{}
	err = json.Unmarshal(body, respStruct)
	if err != nil {
		log.Error(err)
		return err
	}

	if respStruct.Code != 200 {
		return errors.New(respStruct.Error)
	}

	return nil
}

// generate access_token
func GenAccessToken(data string) string {
	h := hmac.New(sha1.New, []byte(sk))
	h.Write([]byte(data))
	signature := base64.URLEncoding.EncodeToString(h.Sum(nil))
	log.Debugf("The signature is  >>>> %s:%s\n", ak, signature)
	return ak + ":" + signature
}
