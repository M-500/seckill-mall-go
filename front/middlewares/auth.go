package middlewares

import "github.com/kataras/iris/v12"

func AuthProduct(ctx iris.Context) {
	uid := ctx.GetCookie("uid")
	if uid == "" {
		ctx.Application().Logger().Debug("用户未登录")
		ctx.Redirect("/user/login")
		return
	}
	ctx.Application().Logger().Debug("登录成功")
	ctx.Next()
}
