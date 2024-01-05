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

type UserController struct {
	Ctx     iris.Context
	Service services.IUserService
	Session *sessions.Session
}

// 实现注册页面的展示
func (u *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

func (u *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (u *UserController) HandleError(ctx iris.Context, err error) {
	if iris.IsErrPath(err) {
		fmt.Println(ctx.Request().RequestURI)
		return // continue.
	}
	ctx.StopWithError(iris.StatusBadRequest, err)
}

// 接收注册页的表单
func (u *UserController) PostRegister() {
	//fmt.Println("擦擦哦")
	nickName := u.Ctx.FormValue("nickName")
	userName := u.Ctx.FormValue("username")
	password := u.Ctx.FormValue("password")
	fmt.Println("哈哈哈", nickName, userName, password)
	user := &models.User{
		NickName: nickName,
		Username: userName,
		Password: password,
	}
	_, err := u.Service.AddUser(user)
	if err != nil {
		u.Ctx.Redirect("/user/error.html")
		//u.Ctx.Application().Logger().Debug(err)
		return
	}
	u.Ctx.Redirect("/user/login")
	return
}

func (c *UserController) PostLogin() mvc.Response {
	//1.获取用户提交的表单信息
	var (
		userName = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)
	//2、验证账号密码正确
	user, isOk := c.Service.IsPwdSuccess(userName, password)
	if !isOk {
		return mvc.Response{
			Path: "/user/login",
		}
	}
	fmt.Println("狗日的", user)
	//3、写入用户ID到cookie中
	//tool.GlobalCookie(c.Ctx, "uid", strconv.FormatInt(user.ID, 10))
	c.Session.Set("userID", strconv.FormatInt(user.ID, 10))

	return mvc.Response{
		Path: "/product/",
	}

}
