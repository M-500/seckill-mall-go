package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int64          `gorm:"primarykey;type bigint;comment:主键自增长ID" json:"id"`
	CreatedAt time.Time      `gorm:"column:create_at;comment:记录创建时间" json:"-"`
	UpdatedAt time.Time      `gorm:"column:update_at;comment:记录更新时间" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

var ModelList = []interface{}{
	&User{},
}
