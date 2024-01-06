package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"seckill-mall-go/models"
	"seckill-mall-go/services"
	"strconv"
)

type ProductController struct {
	Ctx          iris.Context
	ProdService  services.IProdService
	OrderService services.IOrderService
	Session      *sessions.Session
}

// 渲染产品详情页
func (p *ProductController) GetDetail() mvc.View {
	pid := p.Ctx.URLParam("pid")
	num, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	prod, err := p.ProdService.GetProdByID(num)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	fmt.Println(prod)
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": prod,
		},
	}
}

// 渲染抢购结构
func (p *ProductController) GetOrder() mvc.View {
	prodStr := p.Ctx.URLParam("pid")
	uStr := p.Ctx.GetCookie("uid")
	pid, err := strconv.ParseInt(prodStr, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	uid, err := strconv.ParseInt(uStr, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	// 查询商品是否存在
	var orderId int64
	var showMessage string
	product, err := p.ProdService.GetProdByID(pid)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	if product.ProdNum > 0 {
		// 可以开始抢购
		// 1. 扣除商品数量
		product.ProdNum -= 1
		err = p.ProdService.UpdateProd(product) // 商品库存-1 这里会有问题(超卖) 后续会讲解
		if err != nil {
			p.Ctx.Application().Logger().Error(err)
		}
		// 生成订单
		order := &models.Order{
			UserId:      uid,
			ProdId:      pid,
			OrderStatus: models.OrderSuccess,
		}
		orderId, err = p.OrderService.InsertOrder(order)
		if err != nil {
			p.Ctx.Application().Logger().Error(err)
		} else {
			showMessage = "抢购成功！"
		}

	}
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/result.html",
		Data: iris.Map{
			"orderID":     orderId,
			"showMessage": showMessage,
		},
	}
}
