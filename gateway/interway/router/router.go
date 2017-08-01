package router

import (
	c "github.com/goushuyun/taobao-erp/gateway/controller"
	m "github.com/goushuyun/taobao-erp/gateway/middleware"
)

//SetRouterV1 设置seller的router
func SetRouterV1() *m.Router {
	v1 := m.NewWithPrefix("/v1")

	// users
	v1.Register("/users/register", m.Wrap(c.Register))
	v1.Register("/users/check_user_exist", m.Wrap(c.CheckUserExist))

	// sms
	v1.Register("/sms/send_identifying_code", m.Wrap(c.SendIdentifyingCode))

	//book
	v1.Register("/book/get_book_info", m.Wrap(c.GetBookInfo))
	v1.Register("/book/save_book_info", m.Wrap(c.SaveBook))
	v1.Register("/book/update_book_info", m.Wrap(c.UpdateBookInfo))

	// stock
	v1.Register("/stock/get_location_id", m.Wrap(c.GetLocationId))
	v1.Register("/stock/save_single_goods", m.Wrap(c.SaveSingleGoods))
	v1.Register("/stock/save_goods", m.Wrap(c.SaveGoods))
	return v1
}
