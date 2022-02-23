package handler

import (
	"encoding/json"
	"gochicoba/helpers"
	"gochicoba/models"
	"gochicoba/service"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

type ItemHandler struct {
	itemService service.ItemService
}

func NewItemHandler(itemService service.ItemService) ItemHandler {
	return ItemHandler{itemService: itemService}
}

// func ItemContext(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		itemId := chi.URLParam(r, "itemId")
// 		if itemId == "" {
// 			render.Render(w, r, ErrorRenderer(fmt.Errorf("item ID is required")))
// 			return
// 		}
// 		id, err := strconv.Atoi(itemId)
// 		if err != nil {
// 			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid item ID")))
// 		}
// 		ctx := context.WithValue(r.Context(), itemIDKey, id)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// func CreateItem(w http.ResponseWriter, r *http.Request) {
// 	item := &models.Item{}
// 	if err := render.Bind(r, item); err != nil {
// 		render.Render(w, r, ErrBadRequest)
// 		return
// 	}
// 	if err := dbInstance.AddItem(item); err != nil {
// 		render.Render(w, r, ErrorRenderer(err))
// 		return
// 	}
// 	if err := render.Render(w, r, item); err != nil {
// 		render.Render(w, r, ServerErrorRenderer(err))
// 		return
// 	}
// }

func (ih *ItemHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	name := r.URL.Query().Get("name")
	page := r.URL.Query().Get("page")
	view := r.URL.Query().Get("view")

	var list []*models.Item
	var err error
	var start, end *time.Time

	if startDate != "" && endDate != "" {
		s, err := time.Parse("02-01-2006 MST", startDate+" WIB")
		if err != nil {
			helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
			return
		}
		e, err := time.Parse("02-01-2006 15:04:05 999999 MST", endDate+" 23:59:59 999999 WIB")
		if err != nil {
			helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
			return
		}
		// fmt.Println(e)
		start = &s
		end = &e
	}

	var p, v int
	if page == "" {
		page = "0"
	}
	p, err = strconv.Atoi(page)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	if view == "" {
		view = "0"
	}
	v, err = strconv.Atoi(view)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	filter := models.ItemFilter{
		StartDate: start,
		EndDate:   end,
		Name:      name,
		Page:      p,
		View:      v,
	}

	list, err = ih.itemService.GetAllItems(filter)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	response := helpers.BatchItemDate(list)

	helpers.CustomResponse(w, r, http.StatusOK, "success", response)
}

func (ih *ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")
	itemIDInt, err := strconv.Atoi(itemID)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusBadRequest, "id must be integer", err.Error())
		return
	}
	item, err := ih.itemService.GetItem(itemIDInt)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	if item == nil {
		helpers.ErrorResponse(w, r, http.StatusNotFound, "not found", err.Error())
		return
	}
	response := helpers.SingleItemDate(item)
	helpers.CustomResponse(w, r, http.StatusOK, "success", response)
}

func (ih *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item *models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	//fmt.Println(item)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(item)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	itemData, err := ih.itemService.AddItem(item)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	response := helpers.SingleItemDate(itemData)
	helpers.CustomResponse(w, r, http.StatusOK, "success", response)
}

func (ih *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")
	itemIDInt, err := strconv.Atoi(itemID)
	//fmt.Println(itemID)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusBadRequest, "id must be integer", err.Error())
		return
	}
	var item *models.Item
	err = json.NewDecoder(r.Body).Decode(&item)
	//fmt.Println(item)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	item, err = ih.itemService.UpdateItem(itemIDInt, item)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	response := helpers.SingleItemDate(item)
	helpers.CustomResponse(w, r, http.StatusOK, "success", response)
}

func (ih *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")
	itemIDInt, err := strconv.Atoi(itemID)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	err = ih.itemService.DeleteItem(itemIDInt)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", nil)
}

// func DeleteItem(w http.ResponseWriter, r *http.Request) {
// 	itemId := r.Context().Value(itemIDKey).(int)
// 	err := dbInstance.DeleteItem(itemId)
// 	if err != nil {
// 		if err == db.ErrNoMatch {
// 			render.Render(w, r, ErrNotFound)
// 		} else {
// 			render.Render(w, r, ServerErrorRenderer(err))
// 		}
// 		return
// 	}
// }

// func UpdateItem(w http.ResponseWriter, r *http.Request) {
// 	itemId := r.Context().Value(itemIDKey).(int)
// 	itemData := models.Item{}
// 	if err := render.Bind(r, &itemData); err != nil {
// 		render.Render(w, r, ErrBadRequest)
// 		return
// 	}
// 	item, err := dbInstance.UpdateItem(itemId, itemData)
// 	if err != nil {
// 		if err == db.ErrNoMatch {
// 			render.Render(w, r, ErrNotFound)
// 		} else {
// 			render.Render(w, r, ServerErrorRenderer(err))
// 		}
// 		return
// 	}
// 	if err := render.Render(w, r, &item); err != nil {
// 		render.Render(w, r, ServerErrorRenderer(err))
// 		return
// 	}
// }
