package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// func (me *Tag) AutoMigrate() {
// 	db.AutoMigrate(&Tag{})
// }
