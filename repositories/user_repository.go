package repositories

import (
	"database/sql"
	"fmt"
	"seckill-mall-go/common"
	"seckill-mall-go/models"
	"strconv"
)

type IUserRepository interface {
	Conn() error
	SelectByUsername(username string) (user *models.User, err error)
	SelectByID(uid int64) (user *models.User, err error)
	Insert(user *models.User) (uid int64, err error)
}

type UserManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewUserRepository(table string, db *sql.DB) IUserRepository {
	return &UserManager{
		table:     table,
		mysqlConn: db,
	}
}

func (u *UserManager) Conn() error {
	if u.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		u.mysqlConn = mysql
	}
	if u.table == "" {
		u.table = "user"
	}
	return nil
}

func (u *UserManager) SelectByUsername(username string) (user *models.User, err error) {
	if err = u.Conn(); err != nil {
		return
	}
	sqlStr := "SELECT * from " + u.table + " WHERE username = ? "
	row, err := u.mysqlConn.Query(sqlStr, username)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	// 转换结构体
	result := common.GetResultRow(row)
	if len(result) <= 0 {
		return nil, nil
	}
	user = &models.User{}
	common.DataToStructByTagSql(result, user)
	return user, err
}

func (u *UserManager) SelectByID(uid int64) (user *models.User, err error) {
	if err = u.Conn(); err != nil {
		return
	}
	sqlStr := "SELECT * from " + u.table + " WHERE id = " + strconv.FormatInt(uid, 10)
	row, err := u.mysqlConn.Query(sqlStr)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	// 转换结构体
	result := common.GetResultRow(row)
	if len(result) <= 0 {
		return nil, nil
	}
	user = &models.User{}
	common.DataToStructByTagSql(result, user)
	return user, err
}

func (u *UserManager) Insert(user *models.User) (uid int64, err error) {
	if err = u.Conn(); err != nil {
		return
	}
	sqlStr := fmt.Sprintf("INSERT %s SET nick_name=?,username=?,password=?", u.table)
	stmt, err := u.mysqlConn.Prepare(sqlStr)
	if err != nil {
		return
	}
	// 插入数据
	result, err := stmt.Exec(user.NickName,
		user.Username, user.Password,
	)
	if err != nil {
		return
	}
	return result.LastInsertId()
}
