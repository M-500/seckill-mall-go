package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"seckill-mall-go/common"
	"seckill-mall-go/models"
	"seckill-mall-go/services"
	"strconv"
)

type ProdController struct {
	Ctx         iris.Context // 用来获取请求的上下文用的
	ProdService services.IProdService
}

// 获取所有商品
func (p *ProdController) GetAll() mvc.View {
	prods, _ := p.ProdService.GetAllProd()
	return mvc.View{
		Name: "prod/view.html",
		Data: iris.Map{
			"prodList": prods,
		},
	}
}

// 修改商品
func (p *ProdController) PostUpdate() {
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
	p.Ctx.Redirect("/product/all")
}

// 添加商品
func (p *ProdController) GetAdd() mvc.View {
	return mvc.View{
		Name: "prod/add.html",
	}
}

// 商品添加表单
func (p *ProdController) PostAdd() {
	prod := &models.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "seckill"})
	if err := dec.Decode(p.Ctx.Request().Form, prod); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	_, err := p.ProdService.InsertProd(prod)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

// 修改商品页面
func (p *ProdController) GetManager() mvc.View {
	idStr := p.Ctx.URLParam("id")              // 去 URL 中获取商品的Id
	id, err := strconv.ParseInt(idStr, 10, 64) // 转换为int64类型
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	product, err := p.ProdService.GetProdByID(id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	return mvc.View{
		Name: "prod/manager.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

// 删除功能
func (p *ProdController) GetDelete() {
	idStr := p.Ctx.URLParam("id")              // 去 URL 中获取商品的Id
	id, err := strconv.ParseInt(idStr, 10, 64) // 转换为int64类型
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	_, err = p.ProdService.DeleteByID(id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}
