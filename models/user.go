package models

type User struct {
	ID       int64  `sql:"id" json:"id" form:"id"`
	NickName string `sql:"nick_name" json:"nickName" form:"nickName"`
	Username string `sql:"username" json:"username" form:"username"`
	Password string `sql:"password" json:"password" form:"password"`
}
