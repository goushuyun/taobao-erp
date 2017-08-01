package service

import (
	"errors"

	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/misc"

	"golang.org/x/net/context"

	"github.com/goushuyun/taobao-erp/pb"
	"github.com/goushuyun/taobao-erp/stock/db"
	"github.com/wothing/log"
)

type StockServer struct{}

func (s *StockServer) GetLocationId(ctx context.Context, req *pb.Location) (*pb.LocationResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "GetLocationId", "%#v", req))

	err := db.GetLocationId(req)

	if err != nil {
		if err.Error() == "not_found" {
			// create location
			err = db.CreateLocation(req)
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

	return &pb.LocationResp{Code: errs.Ok, Message: "ok", Data: req}, nil
}

func (s *StockServer) SearchLocation(ctx context.Context, req *pb.Location) (*pb.SearchLocationResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SearchGoods", "%#v", req))

	data, err := db.SearchLocation(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}

	return &pb.SearchLocationResp{Code: errs.Ok, Message: "ok", Data: data}, nil
}

// func (s *StockServer) SaveSingleGoods(ctx context.Context, req *pb.Goods) (*pb.NormalResp, error) {
// 	tid := misc.GetTidFromContext(ctx)
// 	defer log.TraceOut(log.TraceIn(tid, "SaveSingleGoods", "%#v", req))
//
// 	// to get book_id
// 	book := &pb.Book{Isbn: req.Isbn}
// 	getBookInfoResp := &pb.BookResp{}
// 	err := misc.CallSVC(ctx, "book", "GetBookInfo", book, getBookInfoResp)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, errs.Wrap(errors.New(err.Error()))
// 	}
// 	book = getBookInfoResp.Data[0]
// 	req.BookId = book.Id
// 	/*
// 		save stock
// 		1. get location_id , create it if not exist
// 		2. get goods_id, create it if not exist
// 		3. update( or create it) map table
// 	*/
// 	var location_id string
//
// 	err = db.GetLocationId(req)
// 	if err.Error() == "no_this_location" {
// 		// create location
// 		loc := &pb.Location{Warehouse: req.Warehouse, Shelf: req.Shelf, Floor: req.Floor, UserId: req.UserId}
// 		err = db.CreateLocation(loc)
// 		if err != nil {
// 			log.Error(err)
// 			return nil, errs.Wrap(errors.New(err.Error()))
// 		}
// 		location_id = loc.LocationId
// 	} else if err != nil {
// 		// met other error
// 		log.Error(err)
// 		return nil, errs.Wrap(errors.New(err.Error()))
// 	} else {
// 		location_id = req.LocationId
// 	}
//
// 	var goods_id string
//
// 	err = db.GetGoodsByBookId(req)
// 	if err.Error() == "not_found" {
// 		// create this goods
// 		err = db.SaveGoods(req)
// 		if err != nil {
// 			log.Error(err)
// 			return nil, errs.Wrap(errors.New(err.Error()))
// 		}
// 		goods_id = req.GoodsId
// 	} else if err != nil {
// 		// met other error
// 		log.Error(err)
// 		return nil, errs.Wrap(errors.New(err.Error()))
// 	} else {
// 		goods_id = req.GoodsId
// 		// update stock
// 		db.UpdateStock(g)
// 	}
//
// 	goods_location_map := &pb.Goods{LocationId: location_id, GoodsId: goods_id, Stock: req.Stock}
//
// 	err = db.SaveMap(goods_location_map)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, errs.Wrap(errors.New(err.Error()))
// 	}
//
// 	return &pb.NormalResp{Code: errs.Ok, Message: "ok"}, nil
// }
