package service

import (
	"github.com/jinzhu/gorm"

	models "github.com/aditiapratama1231/shipping-service/database/models"
	payload "github.com/aditiapratama1231/shipping-service/pkg/request/payload"
)

type ShippingService interface {
	GetShippings() payload.GetShippingResponse
	CreateNewShipping(payload.CreateNewShippingRequest) payload.ShippingResponse
}

type shippingService struct{}

var query *gorm.DB

func NewShippingService(db *gorm.DB) ShippingService {
	query = db
	return shippingService{}
}

func (shippingService) GetShippings() payload.GetShippingResponse {
	var shippings []models.Shipping

	query.Find(&shippings)

	return payload.GetShippingResponse{
		Data: shippings,
	}
}

func (shippingService) CreateNewShipping(data payload.CreateNewShippingRequest) payload.ShippingResponse {
	shipping := models.Shipping{
		ShippingNo:   data.Data.ShippingNo,
		CustomerNote: data.Data.CustomerNote,
		Status:       data.Data.Status,
		ItemID:       data.Data.ItemID,
	}

	err := query.Create(&shipping).Error
	if err != nil {
		return payload.ShippingResponse{
			Message:    "Shipping failed to create : " + err.Error(),
			StatusCode: 500,
			Err:        true,
		}
	}

	return payload.ShippingResponse{
		Message:    "Shipping created successfully",
		Data:       shipping,
		StatusCode: 200,
	}
}