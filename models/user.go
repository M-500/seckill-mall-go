package models

type User struct {
	BaseModel
	NickName string `gorm:"type:string;size:;column:nick_name" seckill:"nickName" sql:"nick_name" json:"nickName" form:"nickName"`
	Username string `gorm:"type:string;size:;column:username" seckill:"username" sql:"username" json:"username" form:"username"`
	Password string `gorm:"type:string;size:;column:password" seckill:"password" sql:"password" json:"password" form:"password"`
}
