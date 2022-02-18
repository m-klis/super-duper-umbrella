package models

import "time"

type Buy struct {
	ID          int       `json:"id"`
	IdUser      int       `json:"id_user"`
	ItemAmount  int       `json:"item_amount"`
	PriceAmount int       `json:"price_amount"`
	CreatedAt   time.Time `json:"create_at"`
}

type BuyResponse struct {
	ID          int    `json:"id"`
	IdUser      int    `json:"id_user"`
	ItemAmount  int    `json:"item_amount"`
	PriceAmount int    `json:"price_amount"`
	CreatedAt   string `json:"create_at"`
}

type BuyResponseService struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	ItemId    int       `json:"item_id"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	Uuid      string    `json:"uuid"`
}

type BuyDetails struct {
	IdBuy       int       `json:"id_buy"`
	IdUser      int       `json:"id_user"`
	IdItem      int       `json:"id_item"`
	ItemName    string    `json:"item_name"`
	ItemPrice   int       `json:"item_price"`
	ItemAmount  int       `json:"item_amount"`
	PriceAmount int       `json:"price_amount"`
	CreatedAt   time.Time `json:"created_at"`
}
