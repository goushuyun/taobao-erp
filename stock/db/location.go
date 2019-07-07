package db

import (
	"github.com/goushuyun/log"
	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
)

func DelLocation(id string) error {
	gs, err := getAllGoodsAtTheLoca(id)
	if err != nil {
		return err
	}

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// if there are goods on the location
	if len(gs) > 0 {
		// 1. delete map relation between goods and location
		delMap := `delete from goods_location_map where location_id = $1`
		if _, err := tx.Exec(delMap, id); err != nil {
			log.Error(err)
			return err
		}

		// 2. update goods stock, reduce the goods stock amount on that location
		for i := 0; i < len(gs); i++ {
			goods := gs[i]

			updateGoodsStock := `update goods set stock = stock - $1 where id = $2`
			log.Infof(`update goods set stock = stock - %d where id = '%s'`, goods.Stock, goods.GoodsId)
			if _, err := tx.Exec(updateGoodsStock, goods.Stock, goods.GoodsId); err != nil {
				log.Error(err)
				return err
			}
		}
	}

	// 3. delete location
	delLocation := `delete from location where id = $1`
	if _, err := tx.Exec(delLocation, id); err != nil {
		log.Error(err)
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func getAllGoodsAtTheLoca(locationID string) (gs []*pb.Goods, err error) {
	query := `select goods_id, stock from goods_location_map where location_id = $1`
	log.Infof(`select goods_id, stock from goods_location_map where location_id = '%s'`, locationID)

	rows, err := DB.Query(query, locationID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for rows.Next() {
		var g = &pb.Goods{}

		if err := rows.Scan(&g.GoodsId, &g.Stock); err != nil {
			log.Error(err)
			return nil, err
		}

		gs = append(gs, g)
	}

	return gs, nil
}
