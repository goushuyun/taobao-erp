package service

import (
	"errors"

	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/misc"

	"golang.org/x/net/context"

	bookDb "github.com/goushuyun/taobao-erp/book/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/goushuyun/taobao-erp/stock/db"
	"github.com/wothing/log"
)

var notFound = "not_found"

func (s *StockServer) SaveGoods(ctx context.Context, req *pb.Goods) (*pb.GoodsResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SaveGoods", "%#v", req))

	err := db.GetGoodsByBookId(req)
	if err != nil {
		if err.Error() == notFound {
			// not found this goods
			err = db.SaveGoods(req)
			if err != nil {
				log.Error(err)
				return nil, errs.Wrap(errors.New(err.Error()))
			}
		} else {
			// met other error
			log.Error(err)
			return nil, errs.Wrap(errors.New(err.Error()))
		}
	}

	return &pb.GoodsResp{Code: errs.Ok, Message: "ok", Data: []*pb.Goods{req}}, nil
}

// search goods from local db
func (s *StockServer) SearchGoods(ctx context.Context, req *pb.GoodsInfo) (*pb.GoodsInfoListResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SearchGoods", "%#v", req))
	models, err, totalCount := db.SearchGoods(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.GoodsInfoListResp{Code: errs.Ok, Message: "ok", Data: models, TotalCount: totalCount}, nil
}

// update the goods info
func (s *StockServer) UpdateGoodsInfo(ctx context.Context, in *pb.Goods) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "UpdateGoodsInfo", "%#v", in))
	err := db.UpdateGoodsInfo(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}

// goods batch upload
func (s *StockServer) GoodsBatchUpload(ctx context.Context, in *pb.GoodsBatchUploadModel) (*pb.GoodsBatchUploadResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GoodsBatchUpload", "%#v", in))

	var successNum, pendingNum int64
	fieldUploadModels := []*pb.GoodsUploadModel{}
	for i := 0; i < len(in.Data); i++ {
		uploadModel := in.Data[i]
		// 1 根绝isbn获取图书信息
		books, err := bookDb.GetBookInfo(&pb.Book{Isbn: uploadModel.Isbn})
		if err != nil {
			log.Error(err)
			fieldUploadModels = append(fieldUploadModels, uploadModel)
			continue
		}
		// handle the book poker situation
		if len(books) > 1 {
			// add new pending check data to db
			err = db.SaveGoodsPendingCheck(&pb.GoodsPendingCheck{Isbn: uploadModel.Isbn, Num: uploadModel.Num, Warehouse: uploadModel.Warehouse, Shelf: uploadModel.Shelf, Floor: uploadModel.Floor, UserId: in.UserId})
			if err != nil {
				log.Error(err)
				fieldUploadModels = append(fieldUploadModels, uploadModel)
				continue
			}
			pendingNum++
			continue
		}

		// 2 根据isbn获取商品信息
		goods := &pb.Goods{BookId: books[0].Id, UserId: in.UserId}
		err = db.GetGoodsByBookId(goods)
		if err != nil {
			if err.Error() == notFound {
				//------------need handle
				err = db.SaveGoods(goods)
				if err != nil {
					log.Error(err)
					fieldUploadModels = append(fieldUploadModels, uploadModel)
					continue
				}
			} else {
				log.Error(err)
				fieldUploadModels = append(fieldUploadModels, uploadModel)
				continue
			}
		}
		// 3 根据位置信息获取位置id
		location := &pb.Location{Warehouse: uploadModel.Warehouse, Floor: uploadModel.Floor, Shelf: uploadModel.Shelf, UserId: in.UserId}
		err = db.GetLocationId(location)
		if err != nil {
			if err.Error() == notFound {
				//------------need handle
				err = db.CreateLocation(location)
				if err != nil {
					log.Error(err)
					fieldUploadModels = append(fieldUploadModels, uploadModel)
					continue
				}
			} else {
				log.Error(err)
				fieldUploadModels = append(fieldUploadModels, uploadModel)
				continue
			}
		}
		// 4 新增或者保存商品和位置的map索引
		maprow := &pb.MapRow{GoodsId: goods.GoodsId, LocationId: location.LocationId, Stock: uploadModel.Num}
		err = db.GetMapRow(maprow)
		if err != nil {
			if err.Error() == notFound {
				_, err = s.SaveMapRow(ctx, maprow)
				if err != nil {
					log.Error(err)
					fieldUploadModels = append(fieldUploadModels, uploadModel)
					continue
				}
			} else {
				log.Error(err)
				fieldUploadModels = append(fieldUploadModels, uploadModel)
				continue
			}
		} else {
			rows := &pb.MapRowBatch{}
			rows.Data = append(rows.Data, maprow)
			_, err = s.UpdateMapRow(ctx, rows)
			if err != nil {
				log.Error(err)
				fieldUploadModels = append(fieldUploadModels, uploadModel)
				continue
			}
		}
		successNum++
	}
	return &pb.GoodsBatchUploadResp{Code: errs.Ok, Message: "ok", SuccessNum: successNum, PendingCheckNum: pendingNum, FailedData: fieldUploadModels}, nil
}

// add a batch upload records
func (s *StockServer) SaveGoodsBatchUploadRecord(ctx context.Context, in *pb.GoodsBatchUploadRecord) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SaveBookBatchUploadRecord", "%#v", in))
	err := db.SaveGoodsBatchUploadRecord(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}

// get the batch upload record list
func (s *StockServer) GetGoodsBatchUploadRecords(ctx context.Context, in *pb.GoodsBatchUploadRecord) (*pb.GoodsBatchUploadRecordsResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetBookBatchUploadRecords", "%#v", in))
	models, err, totalCount := db.GetGoodsBatchUploadRecords(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.GoodsBatchUploadRecordsResp{Code: errs.Ok, Message: "ok", Data: models, TotalCount: totalCount}, nil
}

// get the pending goods check list
func (s *StockServer) GetGoodsPendingCheckList(ctx context.Context, in *pb.GoodsPendingCheck) (*pb.GoodsPendingCheckListResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetGoodsPendingCheckList", "%#v", in))
	models, err, totalCount := db.GetGoodsPendingCheckList(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.GoodsPendingCheckListResp{Code: errs.Ok, Message: "ok", Data: models, TotalCount: totalCount}, nil
}

// deal with the goods check
func (s *StockServer) DealWithGoodsPendingCheckList(ctx context.Context, in *pb.GoodsPendingCheck) (*pb.NormalResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "DealWithGoodsPendingCheckList", "%#v", in))
	err := db.DelGoodsPendingCheckList(in)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
}
