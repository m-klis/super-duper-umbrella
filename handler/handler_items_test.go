package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"gochicoba/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockItemsService struct {
	mock.Mock
}

func (mock *MockItemsService) GetAllItems(it models.ItemFilter) ([]*models.Item, error) {
	args := mock.Called()

	result := args.Get(0)

	return result.([]*models.Item), args.Error(1)
}

func (mock *MockItemsService) GetItem(idItem int) (*models.Item, error) {
	args := mock.Called()

	result := args.Get(0)

	return result.(*models.Item), args.Error(1)
}

func (mock *MockItemsService) AddItem(item *models.Item) (*models.Item, error) {
	args := mock.Called()

	result := args.Get(0)

	return result.(*models.Item), args.Error(1)
}

func (mock *MockItemsService) UpdateItem(idItem int, dataItem *models.Item) (*models.Item, error) {
	args := mock.Called()

	result := args.Get(0)

	return result.(*models.Item), args.Error(1)
}

func (mock *MockItemsService) DeleteItem(idItem int) error {
	args := mock.Called()

	return args.Error(0)
}

func TestGetAllItems(t *testing.T) {
	request, err := http.NewRequest("GET", "/items/", nil)
	response := httptest.NewRecorder()
	items := []*models.Item{
		{
			ID:          1,
			Name:        "TEST1",
			Description: "TEST1",
			CreatedAt:   time.Now(),
			Price:       3000,
		}, {
			ID:          2,
			Name:        "TEST2",
			Description: "TEST2",
			CreatedAt:   time.Time{},
			Price:       4000,
		},
	}

	mockservice := new(MockItemsService)
	mockservice.On("GetAllItems").Return(items, nil)

	h := ItemHandler{itemService: mockservice}
	h.GetAllItems(response, request)

	assert.Equal(t, 200, response.Code)
	assert.NoError(t, err, "Without error")
}

func TestCreateItem(t *testing.T) {
	item := models.Item{
		Name:        "TEST1",
		Description: "TEST1",
		Price:       3000,
	}

	itemRes := models.Item{
		ID:          1,
		Name:        "TEST1",
		Description: "TEST1",
		CreatedAt:   time.Time{},
		Price:       3000,
	}

	byItem, _ := json.Marshal(&item)

	request, err := http.NewRequest("POST", "/users/", bytes.NewBuffer(byItem))
	response := httptest.NewRecorder()

	var itemData *models.Item = &itemRes
	mockService := new(MockItemsService)
	mockService.On("AddItem").Return(itemData, nil)

	h := ItemHandler{itemService: mockService}
	h.CreateItem(response, request)

	assert.Equal(t, 200, response.Code)
	assert.NoError(t, err, "without error")
}

func TestGetItem(t *testing.T) {
	request, err := http.NewRequest("GET", "/items/{itemID}", nil)
	response := httptest.NewRecorder()
	item := models.Item{
		ID:          1,
		Name:        "TEST1",
		Description: "TEST1",
		CreatedAt:   time.Now(),
		Price:       3000,
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("itemID", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	var itemData *models.Item = &item
	mockService := new(MockItemsService)
	mockService.On("GetItem").Return(itemData, nil)

	h := ItemHandler{itemService: mockService}
	h.GetItem(response, request)

	assert.Equal(t, 200, response.Code)
	assert.NoError(t, err, "without error")
}

func TestUpdateItem(t *testing.T) {
	item := models.Item{
		ID:          1,
		Name:        "TEST1",
		Description: "TEST1",
		CreatedAt:   time.Now(),
		Price:       3000,
	}

	byItem, _ := json.Marshal(&item)

	request, err := http.NewRequest("POST", "/items/{itemID}", bytes.NewBuffer(byItem))
	response := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("itemID", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	var itemData *models.Item = &item
	mockService := new(MockItemsService)
	mockService.On("UpdateItem").Return(itemData, nil)

	h := ItemHandler{itemService: mockService}
	h.UpdateItem(response, request)

	assert.Equal(t, 200, response.Code)
	assert.NoError(t, err)
}

func TestDeleteItem(t *testing.T) {
	request, err := http.NewRequest("DELETE", "/items/{itemID}", nil)
	response := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("itemID", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	mockService := new(MockItemsService)
	mockService.On("DeleteItem").Return(nil)

	h := ItemHandler{itemService: mockService}
	h.DeleteItem(response, request)

	assert.Equal(t, 200, response.Code)
	assert.NoError(t, err)
}
