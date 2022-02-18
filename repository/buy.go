package repository

import (
	"gochicoba/models"

	"gorm.io/gorm"
)

type BuyRepository interface {
	GetAllBuys() ([]*models.Buy, error)
	// GetBuy(int) (*models.Buy, error)
	CreateBuy(*models.Buy, []models.BuyDetail) (*models.Buy, error)
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

	err := query.Order("ID ASC").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// func (br *buyRepository) GetBuy(id int) (obd *models.Buy, err error)                    { return }

func (br *buyRepository) CreateBuy(b *models.Buy, bd []models.BuyDetail) (*models.Buy, error) {
	query := br.db.Debug()

	err := query.Create(&b).Error
	if err != nil {
		return nil, err
	}

	// var r *models.Buy
	// err = query.Where("id = ?", b.ID).First(&r).Error
	// if err != nil {
	// 	return nil, err
	// }

	for i := range bd {
		bd[i].IdBuy = b.ID
	}

	err = query.Create(&bd).Error
	if err != nil {
		return nil, err
	}

	return b, nil
}

// func (br *buyRepository) DeleteBuy(id int) (err error)                                  { return }
// func (br *buyRepository) UpdateBuy(id int, bd *models.Buy) (obd *models.Buy, err error) { return }
