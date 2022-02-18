package handler

import (
	"gochicoba/helpers"
	"gochicoba/models"
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

	response := make([]models.BuyResponse, 0)

	for _, l := range list {
		bd := models.BuyResponse{
			ID:          l.ID,
			IdUser:      l.IdUser,
			ItemAmount:  l.ItemAmount,
			PriceAmount: l.PriceAmount,
			CreatedAt:   helpers.ConvertMonth(l.CreatedAt),
		}
		response = append(response, bd)
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", response)
	return
}

// func (bh *BuyHandler) CreateBuy(w http.ResponseWriter, r *http.Request) {
// 	list, err := bh.buyService.CreateBuy()
// 	if err != nil {
// 		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
// 		return
// 	}
// }

// func (bh *BuyHandler) GetBuy(w http.ResponseWriter, r *http.Request) {

// }

// func (bh *BuyHandler) UpdateBuy(w http.ResponseWriter, r *http.Request) {

// }

// func (bh *BuyHandler) DeleteBuy(w http.ResponseWriter, r *http.Request) {

// }
