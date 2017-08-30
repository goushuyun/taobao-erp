package export

import (
	"os"
	"strings"
)

type TaobaoCsvModel struct {
	Product_name          string  `csv:"宝贝名称"`
	Product_category      string  `csv:"宝贝类目"`
	Store_category        string  `csv:"店铺类目"`
	Damage_degree         string  `csv:"新旧程度"`
	Province              string  `csv:"省"`
	City                  string  `csv:"城市"`
	Sale_way              string  `csv:"出售方式"`
	Product_price         float64 `csv:"宝贝价格"`
	Price_floating_range  int64   `csv:"加价幅度"`
	Product_num           int64   `csv:"宝贝数量"`
	Term_of_validity      string  `csv:"有效期"`
	Freight_undertaking   string  `csv:"运费承担"`
	Pingyou_freight       float64 `csv:"平邮"`
	Ems_freight           float64 `csv:"EMS"`
	Express_freight       float64 `csv:"快递"`
	Pay_way               string  `csv:"付款方式"`
	Alipay                string  `csv:"支付宝"`
	Invoice               string  `csv:"发票"`
	Guarantee             string  `csv:"保修"`
	Auto_resend           string  `csv:"自动重发"`
	Upload_warehouse      string  `csv:"放入仓库"`
	Introduce             string  `csv:"橱窗推荐"`
	Deploy_time           string  `csv:"发布时间"`
	Story                 string  `csv:"心情故事"`
	Describe              string  `csv:"宝贝描述"`
	Product_image         string  `csv:"宝贝图片"`
	Product_property      string  `csv:"宝贝属性"`
	Groupon_price         string  `csv:"团购价"`
	Min_groupon_num       string  `csv:"最小团购件数"`
	Delivery_fee_templet  string  `csv:"邮费模版ID"`
	Member_discount       string  `csv:"会员打折"`
	Update_time           string  `csv:"修改时间"`
	Upload_status         string  `csv:"上传状态"`
	Image_status          string  `csv:"图片状态"`
	Rebate                int64   `csv:"返点比例"`
	New_images            string  `csv:"新图片"`
	Video                 string  `csv:"视频"`
	Sale_property_combine string  `csv:"销售属性组合"`
	User_input_id         string  `csv:"用户输入ID串"`
	Product_title_id_map  string  `csv:"用户输入名-值对"`
	Seller_code           string  `csv:"商家编码"`
	Express_weight        string  `csv:"物流重量"`
	Reduce_stock_style    string  `csv:"库存计数"`
}

func PackingTaobaoParam(isbn, category, title, book_title, image, province, city, describe, reduceStockStyle, deliveryFeeTemplete string, stock int64, price, pingyou_fee, ems_fee, express_fee float64) (model *TaobaoCsvModel) {
	model = &TaobaoCsvModel{}
	//自动设置的数据项
	model.Store_category = `",,"`
	model.Damage_degree = "5"
	model.Sale_way = "b"
	model.Price_floating_range = 0
	model.Term_of_validity = "7"
	model.Freight_undertaking = "1"
	model.Pay_way = ""
	model.Alipay = ""
	model.Invoice = "1"
	model.Guarantee = "0"
	model.Auto_resend = "1"
	model.Upload_warehouse = "0"
	model.Introduce = "0"
	model.Deploy_time = ""
	model.Story = ""
	model.Product_image = ""
	model.Product_property = "2043183:2147483647;1636953:2147483647;2045745:4052146;"
	model.Groupon_price = "0"
	model.Min_groupon_num = "0"
	model.Member_discount = "0"
	model.Update_time = ""
	model.Upload_status = "200"
	model.Image_status = "0"
	model.Rebate = 5
	model.Video = ""
	model.Sale_property_combine = ""
	model.User_input_id = `"1636953,2043183"`
	model.Express_weight = ""

	//手动设置的数据项
	title = strings.Replace(title, "\"", `'`, -1)
	model.Product_name = `"` + title + `"`
	model.Product_category = category
	model.Province = province
	model.City = city
	model.Product_price = price
	model.Product_num = stock
	model.Pingyou_freight = pingyou_fee
	model.Ems_freight = ems_fee
	model.Express_freight = express_fee
	describe = strings.Replace(describe, "\"", `'`, -1)
	model.Describe = `"` + describe + `"`
	model.New_images = image
	model.Product_title_id_map = `"` + isbn + "," + book_title + `"`
	model.Seller_code = isbn
	model.Reduce_stock_style = reduceStockStyle
	model.Delivery_fee_templet = deliveryFeeTemplete
	return
}

func SetReadOnly(filepath string) error {
	err := os.Chmod(filepath, 0444)
	return err
}
