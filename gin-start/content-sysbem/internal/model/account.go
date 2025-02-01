package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	// gorm.Model
	ID        int `gorm:"column:id;primaryKey"` // 主键
	Nickname  string
	Password  string
	Email     string
	SuperMan  string
	CreatedAt time.Time      `gorm:"column:created_at"`      // created_at
	UpdatedAt time.Time      `gorm:"column:updated_at"`      // updated_at
	DeletedAt gorm.DeletedAt `gorm:"column:delete_at,index"` // 指定列为索引，适用于软删除
}

func (Account) TableName() string {
	return "account"
}
