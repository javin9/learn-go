package dao

// 里面是最基本的增删改查

import (
	"content-sysbem/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type AccountDao struct {
	db *gorm.DB
}

func NewAccountDao(db *gorm.DB) *AccountDao {
	return &AccountDao{db: db}
}

// 判断是否存在
func (me *AccountDao) IsExist(email string) (bool, error) {
	var account model.Account
	err := me.db.Where(&model.Account{Email: email}).First(&account).Error
	// 没有找到
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	// 报错了
	if err != nil {
		fmt.Printf("AccountDao isExist [%v]", err)
		return false, err
	}
	// 找到了
	return true, nil
}

// 创建
func (me *AccountDao) Create(account *model.Account) error {
	if err := me.db.Create(&account).Error; err != nil {
		return err
	}
	return nil
}

// 通过邮箱查找
func (me *AccountDao) LoginByEmail(email string) (*model.Account, error) {
	var account model.Account
	err := me.db.Where(&model.Account{Email: email}).First(&account).Error
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	// 找到了
	return &account, nil
}
