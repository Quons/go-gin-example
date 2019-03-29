package brand_service

import (
	"github.com/Quons/go-gin-example/models"
	log "github.com/sirupsen/logrus"
)

type Brand struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func (b *Brand) GetBrandList() ([]models.Brand, error) {
	list, err := models.GetBrandList()
	if err != nil {
		log.Error(err)
		return list, err
	}
	return list, nil
}

func (b *Brand) AddBrand() error {
	err := models.AddBrand(&models.Brand{Name: b.Name})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
