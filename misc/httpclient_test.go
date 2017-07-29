package misc

import (
	"testing"

	"github.com/wothing/log"
)

func TestGetReceipt(t *testing.T) {
	url := "https://www.juxinli.com/orgApi/rest/v2/xuexinwang/applications/yiqimei"

	req := map[string]interface{}{
		"selected_website": []interface{}{},
		"basic_info": map[string]interface{}{
			"name":           "武向东",
			"id_card_num":    "410425199004010035",
			"cell_phone_num": "18234134580",
		},
	}

	resp, err := Post(url, req)
	if err != nil {
		log.Error(err)
		return
	}

	log.Debugf("%v", resp)
}
