package controller

import (
	"net/http"

	"github.com/goushuyun/weixin-golang/misc/token"

	"github.com/goushuyun/weixin-golang/misc"
	"github.com/goushuyun/weixin-golang/pb"
)

//AddStore 增加话题
func AddTopic(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Topic{TokenStoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_topic", "AddTopic", req, "title")
}

//DelTopic 删除话题
func DelTopic(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Topic{TokenStoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_topic", "DelTopic", req, "id")
}

//UpdateTopic 更新话题
func UpdateTopic(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Topic{TokenStoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_topic", "UpdateTopic", req, "id")
}

//AddTopicItem 增加话题项
func AddTopicItem(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.TopicItem{}
	misc.CallWithResp(w, r, "bc_topic", "AddTopicItem", req, "goods_id", "topic_id")
}

//DelTopicItem 删除话题项
func DelTopicItem(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.TopicItem{}
	misc.CallWithResp(w, r, "bc_topic", "DelTopicItem", req, "topic_id", "id")
}

//SearchTopics 搜索话题
func SearchTopics(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Topic{TokenStoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_topic", "SearchTopics", req)
}

//SearchTopics 搜索话题
func TopicsInfo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.StoreId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Topic{TokenStoreId: c.StoreId}
	misc.CallWithResp(w, r, "bc_topic", "TopicsInfo", req)
}

//SearchTopics 搜索话题
func TopicsInfoApp(w http.ResponseWriter, r *http.Request) {
	req := &pb.Topic{Status: 1}
	misc.CallWithResp(w, r, "bc_topic", "TopicsInfo", req)
}
