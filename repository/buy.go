package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"gochicoba/models"
	"gochicoba/producer"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type BuyRepository interface {
	GetAllBuys() ([]*models.Buy, error)
	CreateBuy(*models.Buy, []models.BuyDetail) (*models.Buy, error)
	CreateTransaction(float64, *models.RequestTransaction) (*models.Transaction, error)
	CreateTransactionBroker(float64, *models.RequestTransaction) (*models.ResponseTransactionBroker, error)
}

type buyRepository struct {
	db *gorm.DB
	mb *producer.MessageBroker
}

func NewBuyRepository(db *gorm.DB, mb *producer.MessageBroker) BuyRepository {
	return &buyRepository{
		db: db,
		mb: mb,
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

func (br *buyRepository) CreateTransactionBroker(amount float64, rt *models.RequestTransaction) (*models.ResponseTransactionBroker, error) {
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

	_, err = br.mb.Channel.QueueDeclare(
		"QueueService1", // queue name
		true,            // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,
	)
	if err != nil {
		return nil, err
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        req,
	}

	err = br.mb.Channel.Publish(
		"",
		"QueueService1",
		false,
		false,
		message,
	)
	if err != nil {
		return nil, err
	}

	var res *models.ResponseTransactionBroker = &models.ResponseTransactionBroker{
		UserId:    t.UserId,
		Amount:    t.Amount,
		CreatedAt: time.Now(),
		Uuid:      t.Uuid,
		ItemQty:   t.ItemQty,
	}

	return res, nil
}
