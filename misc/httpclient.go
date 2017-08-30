package misc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wothing/log"
)

func Get(url string) ([]byte, error) {
	log.Debug(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("%s: %s", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("%s: %s", url, err)
		return nil, err
	}
	if len(body) > 1000 {
		log.Debugf("resp: %s", body[:1000])
	} else {
		log.Debugf("resp: %s", body)
	}

	return body, nil
}

func GetMap(url string) (map[string]interface{}, error) {
	body, err := Get(url)
	if err != nil {
		log.Errorf("%s: %s", url, err)
		return nil, err
	}

	ret := make(map[string]interface{})
	err = json.Unmarshal(body, &ret)
	if err != nil {
		log.Errorf("%s: %s", body, err)
		return nil, err
	}

	return ret, nil
}

func Post(url string, v interface{}) ([]byte, error) {
	js, err := json.Marshal(v)
	if err != nil {
		log.Errorf("%v: %s", v, err)
		return nil, err
	}
	log.Debugf("post: %s, %s", url, js)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(js))
	if err != nil {
		log.Errorf("%s: %s", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("%s: %s", url, err)
		return nil, err
	}
	if len(body) > 1000 {
		log.Debugf("resp: %s", body[:1000])
	} else {
		log.Debugf("resp: %s", body)
	}

	return body, nil
}

func PostMap(url string, v interface{}) (map[string]interface{}, error) {
	body, err := Post(url, v)
	if err != nil {
		log.Errorf("%s: %s", url, err)
		return nil, err
	}

	ret := make(map[string]interface{})
	err = json.Unmarshal(body, &ret)
	if err != nil {
		log.Errorf("%s: %s", body, err)
		return nil, err
	}

	return ret, nil
}
