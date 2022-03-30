package models

import "time"

type Buy struct {
	ID          int       `json:"id"`
	IdUser      int       `json:"id_user"`
	ItemAmount  int       `json:"item_amount"`
	PriceAmount float64   `json:"price_amount"`
	CreatedAt   time.Time `json:"create_at"`
}

type DataBuy struct {
	ID          int     `json:"id"`
	IdUser      int     `json:"id_user"`
	ItemAmount  int     `json:"item_amount"`
	PriceAmount float64 `json:"price_amount"`
	CreatedAt   string  `json:"create_at"`
}

// response service
// type BuyResponseService struct {
// 	ID        int       `json:"id"`
// 	UserId    int       `json:"user_id"`
// 	ItemId    int       `json:"item_id"`
// 	Price     float64   `json:"price"`
// 	CreatedAt time.Time `json:"created_at"`
// 	// Uuid      string    `json:"uuid"`
// }

type BuyDetail struct {
	IdBuy       int       `json:"id_buy"`
	IdItem      int       `json:"id_item"`
	ItemPrice   float64   `json:"item_price"`
	ItemAmount  int       `json:"item_amount"`
	PriceAmount float64   `json:"price_amount"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateBuy struct {
	DataBuy  DataBuy     `json:"data_buy"`
	DataItem []BuyDetail `json:"data_item"`
}

type RequestTransaction struct {
	UserId  int `json:"user_id" validate:"required"`
	ItemId  int `json:"item_id" validate:"required"`
	ItemQty int `json:"item_qty" validate:"required"`
}

type Transaction struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	ItemId    int       `json:"item_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	Uuid      string    `json:"uuid"`
	ItemQty   int       `json:"item_qty"`
}

type ResponseTransactionBroker struct {
	UserId    int       `json:"user_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	Uuid      string    `json:"uuid"`
	ItemQty   int       `json:"item_qty"`
}

type ResponseTransaction struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       Transaction `json:"data"`
}
