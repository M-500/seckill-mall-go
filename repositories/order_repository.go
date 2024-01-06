package repositories

import (
	"database/sql"
	"fmt"
	"seckill-mall-go/common"
	"seckill-mall-go/models"
	"strconv"
)

// 定义接口
type IOrderRepository interface {
	Conn() error
	Insert(order *models.Order) (ID int64, err error)
	Delete(oid int64) bool
	Update(order *models.Order) error
	SelectByKey(oid int64) (order *models.Order, err error)
	SelectAll() (order []*models.Order, err error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

// 定义结构体
type OrderManagerRepository struct {
	table     string
	mysqlConn *sql.DB
}

func NewOrderRepository(table string, db *sql.DB) IOrderRepository {
	return &OrderManagerRepository{
		table:     table,
		mysqlConn: db,
	}
}

func (o *OrderManagerRepository) Conn() error {
	if o.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}

func (o *OrderManagerRepository) Insert(order *models.Order) (ID int64, err error) {
	if err = o.Conn(); err != nil {
		return
	}
	sqlStr := fmt.Sprintf("INSERT `%s` SET user_id=?,prod_id=?,order_status=?", o.table)
	stmt, err := o.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return
	}
	// 插入数据
	result, err := stmt.Exec(order.UserId,
		order.ProdId, order.OrderStatus,
	)
	if err != nil {
		return
	}
	return result.LastInsertId()
}

func (o *OrderManagerRepository) Delete(oid int64) bool {
	if err := o.Conn(); err != nil {
		return false
	}
	sqlStr := fmt.Sprintf("DELETE FROM %s where id = ?", o.table)
	stmt, err := o.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return false
	}
	_, err = stmt.Exec(oid)
	if err != nil {
		return false
	}
	return true
}

func (o *OrderManagerRepository) Update(order *models.Order) error {
	if err := o.Conn(); err != nil {
		return err
	}
	sqlStr := "UPDATE product SET UserId=?,ProdId=?,OrderStatus=? WHERE id = " + strconv.FormatInt(order.ID, 10)
	stmt, err := o.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	// 更新数据
	_, err = stmt.Exec(order.UserId,
		order.ProdId, order.OrderStatus,
	)
	return nil
}

func (o *OrderManagerRepository) SelectByKey(oid int64) (order *models.Order, err error) {
	if err = o.Conn(); err != nil {
		return
	}
	sqlStr := "SELECT * from " + o.table + " WHERE id = " + strconv.FormatInt(oid, 10)
	row, err := o.mysqlConn.Query(sqlStr)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	// 转换结构体
	result := common.GetResultRow(row)
	if len(result) <= 0 {
		return nil, nil
	}
	order = &models.Order{}
	common.DataToStructByTagSql(result, order)
	return order, err
}

func (o *OrderManagerRepository) SelectAll() ([]*models.Order, error) {
	res := make([]*models.Order, 0)
	if err := o.Conn(); err != nil {
		return res, err
	}
	sqlStr := "SELECT * FROM `order`"
	row, err := o.mysqlConn.Query(sqlStr)
	defer row.Close()
	if err != nil {
		return res, err
	}
	rows := common.GetResultRows(row)
	if len(rows) == 0 {
		return res, err
	}
	for _, v := range rows {
		item := &models.Order{}
		common.DataToStructByTagSql(v, item)
		res = append(res, item)
	}
	return res, nil
}

func (o *OrderManagerRepository) SelectAllWithInfo() (map[int]map[string]string, error) {
	if err := o.Conn(); err != nil {
		return nil, err
	}
	sqlStr := "SELECT o.id ,p.prodName,o.order_status from `order` as o LEFT JOIN product as p ON p.id=o.prod_id"
	rows, err := o.mysqlConn.Query(sqlStr)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return common.GetResultRows(rows), nil
}
