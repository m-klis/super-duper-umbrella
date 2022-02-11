package handler

import (
	"encoding/json"
	"fmt"
	"gochicoba/helpers"
	"gochicoba/models"
	"gochicoba/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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
	list, err := ih.itemService.GetAllItems()
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", list)
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", list)
	return
}

func (ih *ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")

	itemIDInt, err := strconv.Atoi(itemID)
	fmt.Println(itemID)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusBadRequest, "id must be integer", err)
		return
	}
	item, err := ih.itemService.GetItem(itemIDInt)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}
	if item == nil {
		helpers.ErrorResponse(w, r, http.StatusNotFound, "not found", err)
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", item)
	return
}

func (ih *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item *models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	fmt.Println(item)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}

	itemId, err := ih.itemService.AddItem(item)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", itemId)
	return
}

func (ih *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")
	itemIDInt, err := strconv.Atoi(itemID)
	//fmt.Println(itemID)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusBadRequest, "id must be integer", err)
		return
	}
	var item *models.Item
	err = json.NewDecoder(r.Body).Decode(&item)
	//fmt.Println(item)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}

	item, err = ih.itemService.UpdateItem(itemIDInt, item)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}
	helpers.CustomResponse(w, r, http.StatusOK, "success", item)
	return
}

func (ih *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")
	itemIDInt, err := strconv.Atoi(itemID)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}

	err = ih.itemService.DeleteItem(itemIDInt)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", nil)
	return
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
