package service

import (
	"gochicoba/models"
	"gochicoba/repository"
)

type BuyService interface {
	GetAllBuys() ([]*models.Buy, error)
	CreateBuy(models.Buy, []models.BuyDetail) (*models.Buy, error)
	CreateTransaction(float64, *models.RequestTransaction) (*models.Transaction, error)
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
		return nil, err
	}
	return list, nil
}

func (bs *buyService) CreateBuy(db models.Buy, di []models.BuyDetail) (*models.Buy, error) {
	list, err := bs.buyRepo.CreateBuy(&db, di)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (bs *buyService) CreateTransaction(amount float64, rt *models.RequestTransaction) (*models.Transaction, error) {
	res, err := bs.buyRepo.CreateTransaction(amount, rt)
	if err != nil {
		return nil, err
	}
	return res, nil
}
