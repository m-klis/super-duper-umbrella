package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"gochicoba/models"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BuyRepository interface {
	GetAllBuys() ([]*models.Buy, error)
	CreateBuy(*models.Buy, []models.BuyDetail) (*models.Buy, error)
	CreateTransaction(float64, *models.RequestTransaction) (*models.Transaction, error)
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
	query := br.db

	err := query.Order("ID ASC").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (br *buyRepository) CreateBuy(b *models.Buy, bd []models.BuyDetail) (*models.Buy, error) {
	query := br.db

	err := query.Create(&b).Error
	if err != nil {
		return nil, err
	}

	for i := range bd {
		bd[i].IdBuy = b.ID
	}

	err = query.Create(&bd).Error
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (br *buyRepository) CreateTransaction(amount float64, rt *models.RequestTransaction) (*models.Transaction, error) {
	var uid string = uuid.New().String()
	var t *models.Transaction = &models.Transaction{
		UserId:  rt.UserId,
		ItemId:  rt.ItemId,
		Amount:  amount,
		Uuid:    uid,
		ItemQty: rt.ItemQty,
	}

	req, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	bres, err := http.Post("http://localhost:8081/buy/transaction/", "application/json", bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}
	defer bres.Body.Close()
	if bres.StatusCode > 201 {
		return nil, errors.New("service not available")
	}
	body, err := ioutil.ReadAll(bres.Body)
	if err != nil {
		return nil, err
	}

	var res models.ResponseTransaction
	json.Unmarshal([]byte(body), &res)
	return &res.Data, nil
}
