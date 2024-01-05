package services

import (
	"seckill-mall-go/models"
	"seckill-mall-go/repositories"
)

type IOrderService interface {
	GetOrderByID(oid int64) (order *models.Order, err error)
	DeleteOrderByID(oid int64) bool
	UpdateOrder(order *models.Order) error
	InsertOrder(order *models.Order) (ID int64, err error)
	GetAllOrder() (order []*models.Order, err error)
	GetAllOrderWithInfo() (map[int]map[string]string, error)
}

type OrderService struct {
	orderRepo repositories.IOrderRepository
}

func NewOrderService(orderRepo repositories.IOrderRepository) IOrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (o *OrderService) GetOrderByID(oid int64) (order *models.Order, err error) {
	return o.orderRepo.SelectByKey(oid)
}

func (o *OrderService) DeleteOrderByID(oid int64) bool {
	return o.orderRepo.Delete(oid)
}

func (o *OrderService) UpdateOrder(order *models.Order) error {
	return o.orderRepo.Update(order)
}

func (o *OrderService) InsertOrder(order *models.Order) (ID int64, err error) {
	return o.orderRepo.Insert(order)
}

func (o *OrderService) GetAllOrder() (order []*models.Order, err error) {
	return o.orderRepo.SelectAll()
}

func (o *OrderService) GetAllOrderWithInfo() (map[int]map[string]string, error) {
	return o.orderRepo.SelectAllWithInfo()
}
