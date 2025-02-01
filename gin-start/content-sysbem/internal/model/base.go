package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int            `gorm:"column:id;primaryKey"`   // 主键
	CreatedAt time.Time      `gorm:"column:created_at"`      // created_at
	UpdatedAt time.Time      `gorm:"column:updated_at"`      // updated_at
	DeletedAt gorm.DeletedAt `gorm:"column:delete_at,index"` // 指定列为索引，适用于软删除
}
