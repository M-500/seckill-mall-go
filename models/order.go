package models

type Order struct {
	ID          int64 `seckill:"id" sql:"id" json:"id"`
	UserId      int64 `seckill:"userId" sql:"use_id" json:"userId"`
	ProdId      int64 `seckill:"prodId" sql:"prod_id" json:"prodId"`
	OrderStatus int64 `seckill:"orderStatus" sql:"order_status" json:"orderStatus"`
}

const (
	OrderWait = iota
	OrderSuccess
	OrderFailed
)
