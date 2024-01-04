package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"seckill-mall-go/common"
	"seckill-mall-go/models"
	"seckill-mall-go/services"
)

type ProdController struct {
	Ctx         iris.Context // 用来获取请求的上下文用的
	ProdService services.IProdService
}

// 获取所有商品
func (p ProdController) GetAll() mvc.View {
	prods, _ := p.ProdService.GetAllProd()
	return mvc.View{
		Name: "prod/view.html",
		Data: iris.Map{
			"prodList": prods,
		},
	}
}

// 修改商品
func (p ProdController) PostUpdate() {
	product := &models.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "seckill"})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	err := p.ProdService.UpdateProd(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/prod/all")
}
