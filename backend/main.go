package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug") // 设置日志错误等级
	// 注册模板
	template := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
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

	// 注册控制器

	// 启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
