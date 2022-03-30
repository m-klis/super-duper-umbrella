package handler

import (
	"encoding/json"
	"fmt"
	"gochicoba/helpers"
	"gochicoba/models"
	"gochicoba/service"
	"net/http"
	"time"
)

type BuyHandler struct {
	buyService  service.BuyService
	itemService service.ItemService
	userService service.UserService
}

func NewBuyHandler(buyService service.BuyService, itemService service.ItemService, userService service.UserService) BuyHandler {
	return BuyHandler{
		buyService:  buyService,
		itemService: itemService,
		userService: userService,
	}
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

func (bh *BuyHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req *models.RequestTransaction
	var res *models.Transaction
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	user, err := bh.userService.GetUser(req.UserId)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	if user == nil {
		helpers.ErrorResponse(w, r, http.StatusNotFound, "user not found", "")
		return
	}

	item, err := bh.itemService.GetItem(req.ItemId)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	if item == nil {
		helpers.ErrorResponse(w, r, http.StatusNotFound, "item not found", "")
		return
	}

	amount := item.Price * float64(req.ItemQty)

	res, err = bh.buyService.CreateTransaction(amount, req)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	helpers.CustomResponse(w, r, http.StatusOK, "success", res)
}

func (bh *BuyHandler) CreateTransactionBroker(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OKE")

	var req *models.RequestTransaction
	var res *models.ResponseTransactionBroker
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	user, err := bh.userService.GetUser(req.UserId)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	if user == nil {
		helpers.ErrorResponse(w, r, http.StatusNotFound, "user not found", "")
		return
	}

	item, err := bh.itemService.GetItem(req.ItemId)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	if item == nil {
		helpers.ErrorResponse(w, r, http.StatusNotFound, "item not found", "")
		return
	}

	amount := item.Price * float64(req.ItemQty)

	res, err = bh.buyService.CreateTransactionBroker(amount, req)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	helpers.CustomResponse(w, r, http.StatusOK, "success", res)
}
