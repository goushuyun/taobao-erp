package controller

import (
	"net/http"
	"time"

	log "github.com/wothing/log"

	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/stock/export"
	"github.com/tealeg/xlsx"

	"github.com/goushuyun/taobao-erp/misc"
	"github.com/goushuyun/taobao-erp/misc/token"
	"github.com/goushuyun/taobao-erp/pb"
)

func ListGoodsAllLocations(w http.ResponseWriter, r *http.Request) {
	req := &pb.Goods{}

	c := token.Get(r)
	if c != nil && c.UserId != "" {
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "stock", "ListGoodsAllLocations", req, "goods_id")
}

func LocationFazzyQuery(w http.ResponseWriter, r *http.Request) {
	req := &pb.Location{}

	c := token.Get(r)
	if c != nil && c.UserId != "" {
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "stock", "LocationFazzyQuery", req)
}

func UpdateMapRow(w http.ResponseWriter, r *http.Request) {
	req := &pb.MapRowBatch{}

	c := token.Get(r)
	if c != nil && c.UserId != "" {
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "stock", "UpdateMapRow", req, "data")
}

func SaveGoods(w http.ResponseWriter, r *http.Request) {
	req := &pb.Goods{}

	c := token.Get(r)
	if c != nil && c.UserId != "" {
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "stock", "SaveGoods", req, "book_id")
}

func SaveMapRow(w http.ResponseWriter, r *http.Request) {
	req := &pb.MapRow{}

	c := token.Get(r)
	if c != nil && c.UserId != "" {
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "stock", "SaveMapRow", req, "stock", "goods_id", "location_id")
}

func GetLocationId(w http.ResponseWriter, r *http.Request) {
	req := &pb.Location{}

	c := token.Get(r)
	if c != nil && c.UserId != "" {
		req.UserId = c.UserId
	} else {
		misc.RespondMessage(w, r, map[string]interface{}{
			"code":    errs.ErrTokenNotFound,
			"message": "token not found",
		})
		return
	}

	misc.CallWithResp(w, r, "stock", "GetLocationId", req, "warehouse", "shelf", "floor")
}

//get book info
func SearchGoods(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.GoodsInfo{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "SearchGoods", req)
}

// update the goods info
func UpdateGoodsInfo(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}

	req := &pb.Goods{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "UpdateGoodsInfo", req, "goods_id", "remark")
}

// update the goods info
func GoodsBatchUpload(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsBatchUploadModel{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GoodsBatchUpload", req)
}

// add a batch upload record
func SaveGoodsBatchUploadRecord(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsBatchUploadRecord{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "SaveGoodsBatchUploadRecord", req, "origin_file", "origin_filename")
}

// get the batch upload record list
func GetGoodsBatchUploadRecords(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsBatchUploadRecord{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetGoodsBatchUploadRecords", req)
}

// get the pending goods check list
func GetGoodsPendingCheckList(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsPendingCheck{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetGoodsPendingCheckList", req)
}

// deal with the goods check
func DealWithGoodsPendingCheckList(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsPendingCheck{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "DealWithGoodsPendingCheckList", req, "id")
}

// get the location stock
func GetLocationStock(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Location{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetLocationStock", req)
}

// get the location stock
func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Location{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "UpdateLocation", req, "location_id", "warehouse", "shelf", "floor")
}

// get goods pending gathered
func GetGoodsPendingGatherData(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.Goods{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetGoodsPendingGatherData", req)
}

// get goods pending gathered
func GetGoodsShiftRecord(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.GoodsShiftRecord{UserId: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetGoodsShiftRecord", req)
}

// get goods pending gathered
func GetShiftRocordExportDate(w http.ResponseWriter, r *http.Request) {
	c := token.Get(r)
	//检测token
	if c == nil || c.UserId == "" {
		misc.ReturnNotToken(w, r)
		return
	}
	req := &pb.User{Id: c.UserId}
	misc.CallWithResp(w, r, "stock", "GetShiftRocordExportDate", req)
}

// get goods pending gathered
func ExportGoodsShiftRecord(w http.ResponseWriter, r *http.Request) {

	req := &pb.GoodsShiftRecord{}
	body := r.FormValue("params")
	if err := misc.Bytes2Struct([]byte(body), req, "user_id", "start_at", "end_at"); err != nil {
		misc.RespondMessage(w, r, err)
		return
	}

	// Call RPC 请求订单详情
	req.SizeLimit = "none"
	resp, ctx := &pb.GoodsShiftRecordListResp{}, misc.GenContext(r)
	err := misc.CallSVC(ctx, "stock", "GetGoodsShiftRecord", req, resp)
	if err != nil {
		misc.RespondMessage(w, r, err)
		return
	}
	// 开始写入Excel
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("出库单")
	if err != nil {
		log.Error(err)
		res := misc.NewErrResult(errs.ErrInternal, "Error on AddSheet")
		misc.RespondMessage(w, r, res)
		return
	}

	row = sheet.AddRow()
	head := []string{"isbn", "书名", "库位", "架位", "层数", "数量", "日期", "业务类型"}
	row.WriteSlice(&head, len(head))

	for _, item := range resp.Data {
		row = sheet.AddRow()
		var book_no string
		if item.BookNo != "" {
			book_no = "_" + item.BookNo
		}
		row.AddCell().SetString(item.Isbn + book_no)
		row.AddCell().SetString(item.BookTitle)
		row.AddCell().SetString(item.Warehouse)
		row.AddCell().SetString(item.Shelf)
		row.AddCell().SetString(item.Floor)

		row.AddCell().SetInt64(item.Stock)
		row.AddCell().SetString(time.Unix(item.CreateAt, 0).Format("2006-01-02 15:04:05"))
		if item.OperateType == "load" {
			row.AddCell().SetString("入库")
		} else {
			row.AddCell().SetString("出库")
		}
	}

	user := &pb.User{Id: req.UserId, ExportStartAt: req.StartAt, ExportEndAt: req.EndAt}
	misc.CallSVC(ctx, "stock", "UpdateShiftRocordExportDate", user, &pb.NormalResp{})

	filename := "出库单_" + time.Now().Format("2006年01月02日") + ".xlsx"
	w.Header().Set("Content-Disposition",
		`attachment; filename="`+filename+`"; filename*=utf-8''`+filename)

	w.Header().Set("Content-Type",
		`application/vnd.openxmlformats-officedocument.spreadsheetml.sheet; charset=utf-8`)
	file.Write(w)

}

// get goods pending gathered
func ExportCsv(w http.ResponseWriter, r *http.Request) {
	var randFile = time.Now().String() + ".csv"
	w.Header().Set("Content-Type", "application/octet-stream;charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment;filename="+randFile)
	var taobaoModels = make([]*export.TaobaoCsvModel, 0, 1)
	describe := "<P><B><font color=blue>基本信息</font></B></P><P>书名:离散数学</P><P>定价：35.00元</P><P>作者:屈婉玲　等编</P><P>出版社：高等教育出版社</P><P>出版日期：2011-01-01</P><P>ISBN：9787040231250</P><P>字数：550000</P><P>页码：</P><P>版次：1</P><P>装帧：平装</P><P>开本：12k</P><P>商品重量：</P><P><B><font color=blue>编辑推荐</font></B></P><HR SIZE=1><P><p>　　本书特色：<br />　　以教育部计算机科学与技术教学指导委员会制订的计算机科学与技术专业规范为指导，内容涵盖计算机科学技术中常用离散结构的数学基础。<br />　　紧密围绕离散数学的基本概念、基本理论精炼选材，体系严谨，内容丰富；面向计算机科学技术，介绍了很多离散数学在计算机科学技术中的应用。<br />　　强化描述与分析离散结构的基本方法与能力的训练，配有丰富的例题和习题；例题有针对性，分析讲解到位；习题难易结合，适合学生课后练习。<br />　　知识体系采用模块化结构，可以根据不同的教学要求进行调整；语言通俗易懂，深入浅出、突出重点、难点，提示易于出错的地方。<br />　　辅助教学资源丰富，配有用于习题课、包含上千道习题的教学辅导用书《离散数学学习指导与习题解析》，PPT电子教案，教学资源库等。</p></P><P><B><font color=blue>内容提要</font></B></P><HR SIZE=1><P><p>　　本书起源于高等教育出版社1998年出版的《离散数学》，是教育部高等学校“九五”规划教材，2004年作为“十五”规划教材出版了修订版。作为“十一五”规划教材，根据教育部计算机科学与技术专业教学指导委员会提出的《计算机科学与技术专业规范》（CCC2005）的教学要求，本教材对内容进行了较多的调整与更新。<br />　　本书分为数理逻辑、集合论、代数结构、组合数学、图论、初等数论等六个部分。全书既有严谨的、系统的理论阐述，也有丰富的、面向计算机科学技术发展的应用实例，同时选配了大量的典型例题与练习。各章内容按照模块化组织，可以适应不同的教学要求。与本书配套的电子教案和习题辅导用书随后将陆续推出。<br />　　本书可以作为普通高等学校计算机科学与技术专业不同方向的本科生的离散数学教材，也可以供其他专业学生和科技人员阅读参考。</p></P><P><B><font color=blue>目录</font></B></P><HR SIZE=1><P></P><P><B><font color=blue>作者介绍</font></B></P><HR SIZE=1><P><p>　　屈婉玲，1969年毕业于北京大学物理系物理专业，现为北京大学信息科学技术学院教授，博士生导师，中国人工智能学会离散数学专委会委员。主要研究方向是算法设计与分析，发表论文20余篇，出版教材、教学参考书、译著20余本，其中包含多本*规划教材和北京市精品教材。所讲授的离散数学课程被评为国家精品课程，两次被评为北京大学十佳教师，并获得北京市教师称号。曾主持过多项国家教材和课程建设项目，并获得北京市教育教学成果（高等教育）一等奖。</p></P><P><B><font color=blue>序言</font></B></P><HR SIZE=1><P></P>"
	model := export.PackingTaobaoParam("9787040231250", "50000182", "离散数学", "21043943-1_o_2:0:0:;", "北京", "北京", describe, 2, 10, 7, 8, 9)
	log.Debug(model)
	taobaoModels = append(taobaoModels, model)
	//if only ResponseWriter implements the io.Writer interface
	var parser = export.NewCsv(w)
	err := parser.Parse(taobaoModels)
	if err != nil {
		println(err.Error())
	}
}
