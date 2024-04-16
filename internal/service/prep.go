package service

import (
	"fis/internal/entity"
)

const defaultProduct = "1"

var products = map[string]string{
	"Продукт не выбран":         "0",
	"Денежный кредит":           "1",
	"ДК под залог АВТО":         "2",
	"ДК под залог недвижимости": "3",
	"Супер плюс":                "4",
	"Кредитный доктор":          "5",
	"СтопДолг":                  "6",
	"Деньги на карту":           "7",
}

func (s *Service) longApplicationMapping(application *entity.Application) {

}

func (s *Service) fullApplicationToFisMapping(application *entity.Application) {}

func (s *Service) getCreditProduct(productName string) string {
	index, ok := products[productName]
	if !ok {
		return defaultProduct
	}

	return index
}
