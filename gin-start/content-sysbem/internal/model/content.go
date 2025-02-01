package model

import (
	"time"

	"gorm.io/gorm"
)

type Content struct {
	ID             int            `gorm:"column:id;primaryKey"`   // 主键
	Title          string         `gorm:"column:title"`           // 标题
	Description    string         `gorm:"column:description"`     // 描述
	Author         string         `gorm:"column:author"`          // 作者
	Thumbnail      string         `gorm:"column:thumbnail"`       // 封面
	Category       int            `gorm:"column:category"`        // 分类
	Duration       int            `gorm:"column:duration"`        // 时长
	Resolution     string         `gorm:"column:resolution"`      // 分辨率 例如： 720p 1080p
	FileSize       int64          `gorm:"column:file_size"`       // 文件大小 单位：字节
	Format         string         `gorm:"column:format"`          // 文件格式 MP4、AVI
	Quality        int            `gorm:"column:quality"`         // 视频质量 1 高清，2标清
	ApprovalStatus int            `gorm:"column:approval_status"` // 审核状态 1 审核中 2 审核通过 3 审核不通过
	CreatedAt      time.Time      `gorm:"column:created_at"`      // created_at
	UpdatedAt      time.Time      `gorm:"column:updated_at"`      // updated_at
	DeletedAt      gorm.DeletedAt `gorm:"column:delete_at,index"` // 指定列为索引，适用于软删除
}

func (Content) TableName() string {
	return "content"
}
