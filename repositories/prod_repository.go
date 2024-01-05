package repositories

import (
	"database/sql"
	"seckill-mall-go/common"
	"seckill-mall-go/models"
	"strconv"
)

// 开发流程
// 1. 开发接口
// 2. 实现接口

type IProduct interface {
	// 连接数据库
	Conn() error
	// 插入数据接口
	Insert(product *models.Product) (int64, error)
	// 删除
	Delete(int64) (bool, error)
	// 更新
	Update(product *models.Product) error
	// 查询
	SelectByKey(int64) (*models.Product, error)
	// 查找所有产品
	SelectAll() ([]*models.Product, error)
}

type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductManager(tb string, db *sql.DB) IProduct {
	return &ProductManager{
		table:     tb,
		mysqlConn: db,
	}
}

// 数据库连接
func (p *ProductManager) Conn() error {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return nil
}
func (p *ProductManager) Insert(product *models.Product) (pid int64, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sqlStr := "INSERT product SET prodName=?,prodNum=?,prodCover=?,prodImages=?,prodUrl=?"
	stmt, err := p.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return
	}
	// 插入数据
	result, err := stmt.Exec(product.ProdName,
		product.ProdNum, product.ProdCover, product.ProdImages, product.ProdUrl,
	)
	if err != nil {
		return
	}
	return result.LastInsertId()
}
func (p *ProductManager) Delete(pid int64) (exist bool, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sqlStr := "DELETE FROM product where id = ?"
	stmt, err := p.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return
	}
	_, err = stmt.Exec(pid)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (p *ProductManager) Update(product *models.Product) (err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sqlStr := "UPDATE product SET prodName=?,prodNum=?,prodCover=?,prodImages=?,prodUrl=? WHERE id = " + strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return
	}
	// 更新数据
	_, err = stmt.Exec(product.ProdName,
		product.ProdNum, product.ProdCover, product.ProdImages, product.ProdUrl,
	)
	return
}

// 根据名字匹配
func (p *ProductManager) SelectByKey(pid int64) (product *models.Product, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sqlStr := "SELECT * from " + p.table + " WHERE id = " + strconv.FormatInt(pid, 10)
	row, err := p.mysqlConn.Query(sqlStr)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	// 转换结构体
	result := common.GetResultRow(row)
	if len(result) <= 0 {
		return nil, nil
	}
	product = &models.Product{}
	common.DataToStructByTagSql(result, product)
	return product, err
}
func (p *ProductManager) SelectAll() (product []*models.Product, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sqlStr := "SELECT * from " + p.table
	row, err := p.mysqlConn.Query(sqlStr)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	rows := common.GetResultRows(row)
	if len(rows) == 0 {
		return
	}
	for _, v := range rows {
		item := &models.Product{}
		common.DataToStructByTagSql(v, item)
		product = append(product, item)
	}
	return product, nil
}
