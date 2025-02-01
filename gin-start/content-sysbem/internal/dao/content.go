package dao

import (
	"content-sysbem/internal/model"
	"fmt"

	"gorm.io/gorm"
)

// 里面是最基本的增删改查
type ContentDao struct {
	db *gorm.DB
}

func NewContentDao(db *gorm.DB) *ContentDao {
	return &ContentDao{db: db}
}

// 创建
func (me *ContentDao) Create(content *model.Content) error {
	if err := me.db.Create(&content).Error; err != nil {
		fmt.Printf("ContentDao Create error: %v", err)
		return err
	}
	return nil
}

// 判断是否存在
func (me *ContentDao) ExistByID(id int) (bool, error) {
	var content model.Content
	err := me.db.Select("id").Where("id = ?", id).First(&content).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		fmt.Printf("ContentDao ExistByID error: %v", err)
		return false, err
	}
	return true, nil
}

// 更新
func (me *ContentDao) Update(id int, content *model.Content) error {
	if err := me.db.Where("id = ?", id).Updates(&content).Error; err != nil {
		fmt.Printf("ContentDao Update error: %v", err)
		return err
	}
	return nil
}

// 删除 软删除 https://gorm.io/zh_CN/docs/delete.html
func (me *ContentDao) DeleteByID(id int) error {
	if err := me.db.Where("id = ?", id).Delete(&model.Content{}).Error; err != nil {
		fmt.Printf("ContentDao DeleteByID error: %v", err)
		return err
	}
	return nil
}

// 查询
type FindParams struct {
	ID       int
	Page     int
	PageSize int
	Title    string
	Author   string
}

func (me *ContentDao) Find(params *FindParams) ([]model.Content, int, error) {

	// 指定模型
	query := me.db.Model(&model.Content{})
	// 查询条件
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	// 模糊查询
	if params.Title != "" {
		query = query.Where("title like ?", "%"+params.Title+"%")
	}
	if params.Author != "" {
		query = query.Where("author like ?", "%"+params.Author+"%")
	}

	// 分页查询 总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		fmt.Printf("ContentDao Find error: %v", err)
		return nil, 0, err
	}

	// 计算偏移量
	var page, page_size = 1, 10
	if params.Page > 0 {
		page = params.Page
	}

	if params.PageSize > 0 {
		page_size = params.PageSize
	}

	offset := (page - 1) * page_size
	var contentList []model.Content
	if err := query.Offset(offset).Limit(page_size).Find(&contentList).Error; err != nil {
		fmt.Printf("ContentDao Find error: %v", err)
		return nil, 0, err
	}
	return contentList, 0, nil
}
