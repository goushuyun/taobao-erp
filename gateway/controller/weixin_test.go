package controller

import (
	"encoding/xml"
	"testing"

	"github.com/wothing/log"
)

func TestParseXML(t *testing.T) {
	var text = `<xml>
    <ToUserName><![CDATA[gh_d86423f306aa]]></ToUserName>
    <Encrypt><![CDATA[3vU/6JtPfXupWx5SYS6OC4wpLAWu5vhw+DBku5P9lFgZedlJJC5miOrPZsg6pAXOT12KgnmJSsTs8yvBmkhI5Cp63OUW24hu7ghl4cLdJ1RfXIU1zoJtKa1ZeLdeqF3iHm0cmAREzv2SAAVa/kcTKWj8T+oFENUDYyEO6i5kMlP3t62Ef4Je5m+sLwY9kDKKoFjmQmPbGTh5PbTakzuuzDGUDK3p5OiIc2XgjuDO0qdWPvIU72fw8MoNL+2+WJCBYAP0AFIyJzMo5dHMPBsOCdFRRKVuOcyYL8k1KDZ/FDQ2lRa64u2okxNiaWFStZQQMf8TmlXsnE0g01GfOaAJMBX4HaArxlaHxXBlhGS5ZHB1NxcnNZ+YPzAOMuVeHq4bZk1j8DkRyRqzdQ7V9Bk5C17QmY8Gn//VjiYME7lc93NbE141ahYO81Jkv8IznZIs8iXHXFFiPdCk79dG+QJbYA==]]></Encrypt>
</xml>`

	type callback struct {
		XMLName    xml.Name `xml:"xml"`
		ToUserName string   `xml:"ToUserName"`
		Encrypt    string   `xml:"Encrypt"`
	}

	cb := &callback{}
	err := xml.Unmarshal([]byte(text), cb)
	if err != nil {
		t.Fatal(err)
	}

	c, err := getCrypter()
	if err != nil {
		t.Fatal(err)
	}

	crypter := *c

	crypterText, _, err := crypter.Decrypt(cb.Encrypt)
	if err != nil {
		t.Fatal(err)
	}

	log.Debugf("解密后的文本： %s\n", crypterText)
}
