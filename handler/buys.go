package handler

import (
	"encoding/json"
	"gochicoba/helpers"
	"gochicoba/models"
	"gochicoba/service"
	"net/http"
	"time"
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

	response := make([]models.DataBuy, 0)

	for _, l := range list {
		bd := models.DataBuy{
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

func (bh *BuyHandler) CreateBuy(w http.ResponseWriter, r *http.Request) {
	var cb *models.CreateBuy
	err := json.NewDecoder(r.Body).Decode(&cb)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	var layoutFormat string = "2006-01-02 15:04:05"
	createdAt, _ := time.Parse(layoutFormat, cb.DataBuy.CreatedAt)
	var db models.Buy = models.Buy{
		ID:          cb.DataBuy.ID,
		IdUser:      cb.DataBuy.IdUser,
		ItemAmount:  cb.DataBuy.ItemAmount,
		PriceAmount: cb.DataBuy.PriceAmount,
		CreatedAt:   createdAt,
	}
	var di []models.BuyDetail
	di = append(di, cb.DataItem...)

	// db *models.DataBuy, di []*models.DataItem
	data, err := bh.buyService.CreateBuy(db, di)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	helpers.CustomResponse(w, r, http.StatusOK, "success", data)
}

// func (bh *BuyHandler) GetBuy(w http.ResponseWriter, r *http.Request) {

// }

// func (bh *BuyHandler) UpdateBuy(w http.ResponseWriter, r *http.Request) {

// }

// func (bh *BuyHandler) DeleteBuy(w http.ResponseWriter, r *http.Request) {

// }
