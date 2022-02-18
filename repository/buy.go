package repository

import (
	"gochicoba/models"

	"gorm.io/gorm"
)

type BuyRepository interface {
	GetAllBuys() ([]*models.Buy, error)
	// GetBuy(int) (*models.Buy, error)
	// CreateBuy(*models.Buy) (*models.Buy, error)
	// DeleteBuy(int) error
	// UpdateBuy(int, *models.Buy) (*models.Buy, error)
}

type buyRepository struct {
	db *gorm.DB
}

func NewBuyRepository(db *gorm.DB) BuyRepository {
	return &buyRepository{
		db: db,
	}
}

func (br *buyRepository) GetAllBuys() ([]*models.Buy, error) {
	var list []*models.Buy
	query := br.db.Debug()
	// query := ir.db

	err := query.Order("ID ASC").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// func GetBuy(id int) (obd *models.Buy, err error)                    { return }

// func CreateBuy(bd *models.Buy) (obd *models.Buy, err error) { return }

// func DeleteBuy(id int) (err error)                                  { return }
// func UpdateBuy(id int, bd *models.Buy) (obd *models.Buy, err error) { return }
