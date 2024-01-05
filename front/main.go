package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"seckill-mall-go/common"
	"seckill-mall-go/front/web/controllers"
	"seckill-mall-go/repositories"
	"seckill-mall-go/services"
	"time"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug") // 设置日志错误等级
	// 注册模板
	template := iris.HTML("front/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	// 设置静态文件目录
	//app.StaticContent("/assets", "./backend/web/assets")
	app.HandleDir("/public", "./front/web/public")
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

	sess := sessions.New(sessions.Config{
		Cookie:  "seckill-mall-go",
		Expires: 60 * time.Minute,
	})
	// 1. 注册User控制器
	userRepo := repositories.NewUserRepository("product", db)
	userService := services.NewUserService(userRepo)
	userParty := app.Party("/user")
	user := mvc.New(userParty)
	user.Register(ctx, userService, sess.Start)
	user.Handle(new(controllers.UserController))

	app.Run(
		iris.Addr("0.0.0.0:8081"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
