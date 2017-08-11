package register

import (
	"fmt"

	"github.com/goushuyun/taobao-erp/book/service"
	"github.com/robfig/cron"
)

//注册时间轮询
func RegisterBookPolling(cron *cron.Cron) {

	one_minute_spec := "0 0/1 * * * *"
	two_minute_spec := "0 0/2 * * * *"
	three_minute_spec := "0 0/3 * * * *"
	five_minute_spec := "0 0/5 * * * *"
	every_day_on_before_dawn_1 := "0 0 1 * * *" //时间轮询表达式 每天凌晨1：00执行一次
	fmt.Println(one_minute_spec, two_minute_spec, three_minute_spec, five_minute_spec, every_day_on_before_dawn_1)
	go func() {
		service.HandlePendingBook()
	}()
	// 增加 时间 轮训
	cron.AddFunc(every_day_on_before_dawn_1, service.GetAllIncompleteBook)

}
