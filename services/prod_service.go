package services

import (
	"seckill-mall-go/models"
	"seckill-mall-go/repositories"
)

type IProdService interface {
	GetProdByID(pid int64) (prod *models.Product, err error)
	GetAllProd() (prods []*models.Product, err error)
	DeleteByID(pid int64) (exist bool, err error)
	InsertProd(product *models.Product) (pid int64, err error)
	UpdateProd(product *models.Product) (err error)
}

type ProdServiceManager struct {
	prodRepo repositories.IProduct
}

func NewProdServiceManager(prodRepo repositories.IProduct) IProdService {
	return &ProdServiceManager{
		prodRepo: prodRepo,
	}
}

func (p *ProdServiceManager) GetProdByID(pid int64) (prod *models.Product, err error) {
	return p.prodRepo.SelectByKey(pid)
}

func (p *ProdServiceManager) GetAllProd() (prods []*models.Product, err error) {
	return p.prodRepo.SelectAll()
}

func (p *ProdServiceManager) DeleteByID(pid int64) (exist bool, err error) {
	return p.prodRepo.Delete(pid)
}

func (p *ProdServiceManager) InsertProd(product *models.Product) (pid int64, err error) {
	return p.prodRepo.Insert(product)
}

func (p *ProdServiceManager) UpdateProd(product *models.Product) (err error) {
	return p.prodRepo.Update(product)
}
