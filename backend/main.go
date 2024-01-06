package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"seckill-mall-go/backend/web/controllers"
	"seckill-mall-go/common"
	"seckill-mall-go/repositories"
	"seckill-mall-go/services"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug") // 设置日志错误等级
	// 注册模板
	template := iris.HTML("backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	// 设置静态文件目录
	//app.StaticContent("/assets", "./backend/web/assets")
	app.HandleDir("/assets", "./backend/web/assets")
	// 处理异常跳转指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错啦！"))
		ctx.ViewLayout("")
		err := ctx.View("shared/error.html")
		if err != nil {
			return
		}
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 1. 连接MySQL
	db, err := common.NewMysqlConn()
	if err != nil {
		panic("数据库连接失败")
	}

	// 注册控制器
	prodRepo := repositories.NewProductRepository("product", db)
	prodService := services.NewProdServiceManager(prodRepo)
	prodParty := app.Party("/product")
	prod := mvc.New(prodParty)
	prod.Register(ctx, prodService)
	prod.Handle(new(controllers.ProdController))

	// 注册订单控制器
	orderRepo := repositories.NewOrderRepository("order", db)
	orderService := services.NewOrderService(orderRepo)
	orderParty := app.Party("/order")
	order := mvc.New(orderParty)
	order.Register(ctx, orderService)
	order.Handle(new(controllers.OrderController))
	// 启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
