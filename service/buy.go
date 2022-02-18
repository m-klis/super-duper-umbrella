package service

import (
	"gochicoba/models"
	"gochicoba/repository"
)

type BuyService interface {
	GetAllBuys() ([]*models.Buy, error)
	// GetBuy(int) (*models.Buy, error)
	CreateBuy(models.Buy, []models.BuyDetail) (*models.Buy, error)
	// DeleteBuy(int) error
	// UpdateBuy(int, *models.Item) (*models.Item, error)
}

type buyService struct {
	buyRepo repository.BuyRepository
}

func NewBuyService(buyRepo repository.BuyRepository) BuyService {
	return &buyService{
		buyRepo: buyRepo,
	}
}

func (bs *buyService) GetAllBuys() ([]*models.Buy, error) {
	list, err := bs.buyRepo.GetAllBuys()
	if err != nil {
		return list, err
	}
	return list, err
}

// func (bs *buyService) GetBuy(id int) (obd *models.Buy, err error) {
// 	return nil, nil
// }

func (bs *buyService) CreateBuy(db models.Buy, di []models.BuyDetail) (*models.Buy, error) {
	list, err := bs.buyRepo.CreateBuy(&db, di)
	if err != nil {
		return list, err
	}
	return list, err
}

// func (bs *buyService) DeleteBuy(id int) (err error) {
// 	return err
// }
// func (bs *buyService) UpdateBuy(id int, bd *models.Buy) (obd *models.Buy, err error) {
// 	return nil, nil
// }
