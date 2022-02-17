package handler

import (
	"gochicoba/helpers"
	"gochicoba/service"
	"net/http"
)

type BuyHandler struct {
	buyService service.BuyService
}

func NewBuyHandler(buyService service.BuyService) BuyHandler {
	return BuyHandler{buyService: buyService}
}

func (bh *BuyHandler) GetAllBuys(w http.ResponseWriter, r *http.Request) {
	list, err := bh.buyService.GetAllBuys()
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	helpers.CustomResponse(w, r, http.StatusOK, "success", list)
	return
}

// func (bh *BuyHandler) CreateBuy(w http.ResponseWriter, r *http.Request) {

// }

// func (bh *BuyHandler) GetBuy(w http.ResponseWriter, r *http.Request) {

// }

// func (bh *BuyHandler) UpdateBuy(w http.ResponseWriter, r *http.Request) {

// }

// func (bh *BuyHandler) DeleteBuy(w http.ResponseWriter, r *http.Request) {

// }
