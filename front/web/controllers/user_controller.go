package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"seckill-mall-go/common"
	"seckill-mall-go/models"
	"seckill-mall-go/services"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
	Session *sessions.Session
}

// 实现注册页面的展示
func (u *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

// 接收注册页的表单
func (u *UserController) PostRegister() {

	//fmt.Println("擦擦哦")
	//nickName := u.Ctx.FormValue("nickName")
	//userName := u.Ctx.FormValue("username")
	//password := u.Ctx.FormValue("password")
	//fmt.Println("哈哈哈", nickName, userName, password)
	//user := &models.User{
	//	NickName: nickName,
	//	Username: userName,
	//	Password: password,
	//}
	user := &models.User{}
	err := u.Ctx.Request().ParseForm()
	if err != nil {
		fmt.Println(err.Error(), "妈的")
		return
	}
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "form"})
	if err := dec.Decode(u.Ctx.Request().Form, user); err != nil {
		u.Ctx.Redirect("/user/error.html")
		return
	}
	_, err = u.Service.AddUser(user)
	if err != nil {
		u.Ctx.Redirect("/user/error.html")
		return
	}
	u.Ctx.Redirect("/user/login.html")
}
