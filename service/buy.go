package service

import (
	"gochicoba/models"
	"gochicoba/repository"
)

type BuyService interface {
	GetAllBuys() ([]*models.Buy, error)
	// GetBuy(int) (*models.Buy, error)
	// AddBuy(*models.Buy) (*models.Buy, error)
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

// func GetBuy(id int) (obd *models.Buy, err error) {
// 	return nil, nil
// }
// func AddBuy(bd *models.Buy) (obd *models.Buy, err error) {
// 	return nil, nil
// }
// func DeleteBuy(id int) (err error) {
// 	return err
// }
// func UpdateBuy(id int, bd *models.Buy) (obd *models.Buy, err error) {
// 	return nil, nil
// }
