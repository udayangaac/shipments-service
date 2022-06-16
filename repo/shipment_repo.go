package repo

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/udayangaac/shipments-service/repo/entity"
)

// UserRepo handles transactions related to the user.
type ShipmentRepo interface {
	// Save save shipment to the database.
	Save(ctx context.Context, shipment entity.Shipment) (err error)
}

// NewShipmentRepo create an instance of ShipmentRepo implementation.
func NewShipmentRepo(db *gorm.DB) ShipmentRepo {
	return &shipmentRepo{
		DB: db,
	}
}

type shipmentRepo struct {
	DB *gorm.DB
}

// Save save shipment to the database.
func (s *shipmentRepo) Save(ctx context.Context, shipment entity.Shipment) (err error) {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&shipment.Consignee).Error; err != nil {
			return err
		}
		if err := tx.Create(&shipment.Shipper).Error; err != nil {
			return err
		}

		if err := tx.Create(&shipment.Invoice).Error; err != nil {
			return err
		}

		if err := tx.Create(&shipment).Error; err != nil {
			return err
		}

		return nil
	})
}
