package models

import (
	"github.com/jinzhu/gorm"
)

type Brand struct {
	ID   int64  `gorm:"primary_key;column:id"`
	Name string `gorm:"column:name"`
}

func GetBrandList() ([]Brand, error) {
	var brandList []Brand
	err := readDB().Find(&brandList).Error
	if err == gorm.ErrRecordNotFound {
		return brandList, nil
	}
	return brandList, err
}


func AddBrand(brand *Brand) error {
	if err := WriteDB().Create(brand).Error; err != nil {
		return err
	}
	return nil
}
