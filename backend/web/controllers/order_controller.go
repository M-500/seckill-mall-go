package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"seckill-mall-go/common"
	"seckill-mall-go/models"
	"seckill-mall-go/services"
)

type OrderController struct {
	Ctx          iris.Context // 用来获取请求的上下文用的
	OrderService services.IOrderService
}

// 查询所有订单相关信息
func (o *OrderController) Get() mvc.View {
	order, err := o.OrderService.GetAllOrderWithInfo()
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	for _, i2 := range order {
		fmt.Println(i2)
	}
	return mvc.View{
		Name: "order/view.html",
		Data: iris.Map{
			"order": order,
		},
	}

}

// 修改订单
func (o *OrderController) PostUpdate() {
	order := &models.Order{}
	o.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "seckill"})
	if err := dec.Decode(o.Ctx.Request().Form, order); err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	err := o.OrderService.UpdateOrder(order)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	o.Ctx.Redirect("/order/all")
}

//// 新增订单
//func (p *OrderController) GetAdd() mvc.View {
//	return mvc.View{
//		Name: "prod/add.html",
//	}
//}
