package models

import "time"

type Buy struct {
	ID          int       `json:"id"`
	IdUser      int       `json:"id_user"`
	IdItem      int       `json:"id_item"`
	IdDetails   int       `json:"id_details"`
	ItemQty     int       `json:"item_qty"`
	ItemAmount  int       `json:"item_amount"`
	PriceAmount int       `json:"price_amount"`
	CreatedAt   time.Time `json:"create_at"`
}

type BuyResponse struct {
	ID          int       `json:"id"`
	IdUser      int       `json:"id_user"`
	IdItem      int       `json:"id_item"`
	IdDetails   int       `json:"id_details"`
	ItemQty     int       `json:"item_qty"`
	ItemAmount  int       `json:"item_amount"`
	PriceAmount int       `json:"price_amount"`
	CreatedAt   time.Time `json:"create_at"`
}

type BuyDetails struct {
	ID          int       `json:"id"`
	IdUser      int       `json:"id_user"`
	IdItem      int       `json:"id_item"`
	ItemName    string    `json:"item_name"`
	ItemPrice   int       `json:"item_price"`
	ItemAmount  int       `json:"item_amount"`
	PriceAmount int       `json:"price_amount"`
	CreatedAt   time.Time `json:"created_at"`
}
