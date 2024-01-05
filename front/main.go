package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"seckill-mall-go/common"
	"seckill-mall-go/front/web/controllers"
	"seckill-mall-go/models"
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
	dsn := "admin:123456@tcp(192.168.1.52:3306)/seckill_mall?charset=utf8mb4&parseTime=True&loc=Local"
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢查询阈值
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false, // 禁止彩色打印
		},
	)
	gormConf := &gorm.Config{
		Logger: gormLogger,
	}
	err := common.OpenDB(dsn, gormConf, 10, 20, models.ModelList...)
	if err != nil {
		panic(any(err))
		return
	}

	sess := sessions.New(sessions.Config{
		Cookie:  "seckill-mall-go",
		Expires: 60 * time.Minute,
	})
	// 1. 注册User控制器
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userParty := app.Party("/user")
	user := mvc.New(userParty)
	user.Register(ctx, userService, sess.Start)
	user.Handle(new(controllers.UserController))

	app.Run(
		iris.Addr("127.0.0.1:8085"),
		//iris.WithoutServerError(iris.ErrServerClosed),
		//iris.WithOptimizations,
	)
}
