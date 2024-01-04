package models

type Product struct {
	ID         int64  `seckill:"id" sql:"id" json:"id"`
	ProdName   string `seckill:"prodName" sql:"prodName" json:"prodName"`
	ProdNum    int64  `seckill:"prodNum" sql:"prodNum" json:"prodNum"`
	ProdCover  string `seckill:"prodCover" sql:"prodCover" json:"prodCover"`
	ProdImages string `seckill:"prodImages" sql:"prodImages" json:"prodImages"`
	ProdUrl    string `seckill:"prodUrl" sql:"prodUrl" json:"prodUrl"`
}
