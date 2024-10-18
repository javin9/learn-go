package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name  string `json:"name"`
	State int    `json:"state"`
}

func ExistTagName(name string) (bool, error) {
	var tag = Tag{}
	result := db.Where(&Tag{Name: name}).First(&tag)
	// 查询的时候报错
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}

	// 查询的时候正确，但是没结果
	if tag.ID == 0 {
		return false, nil
	}

	// 查询的时候正确，有结果
	return true, nil
}

// add
func AddTag(name string, state int) (uint, error) {
	exist, err := ExistTagName(name)
	if err != nil {
		return 0, err
	}

	if exist {
		return 0, fmt.Errorf("%s is exist", name)
	}
	tag := Tag{Name: name, State: state}
	if err := db.Create(&tag).Error; err != nil {
		return 0, err
	}

	return tag.ID, nil
}

// delete
func DeleteTag(id uint) error {
	if result := db.Delete(&Tag{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func EditTag(id uint, data any) error {
	result := db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
