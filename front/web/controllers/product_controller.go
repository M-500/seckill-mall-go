package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"seckill-mall-go/services"
	"strconv"
)

type ProductController struct {
	Ctx         iris.Context
	ProdService services.IProdService
	Session     *sessions.Session
}

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
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": prod,
		},
	}
}
