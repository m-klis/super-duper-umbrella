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
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

type MockUsersService struct {
	mock.Mock
}

func (mock *MockUsersService) GetAllUsers(models.UserFilter) ([]*models.User, error) {
	args := mock.Called()

	result := args.Get(0)

	return result.([]*models.User), args.Error(1)
}

func (mock *MockUsersService) GetUser(userId int) (userData *models.User, err error) {
	args := mock.Called()

	result := args.Get(0)

	return result.(*models.User), args.Error(1)
}

func (mock *MockUsersService) AddUser(user *models.User) (*models.User, error) {
	args := mock.Called()

	result := args.Get(0)

	return result.(*models.User), args.Error(1)
}

func (mock *MockUsersService) DeleteUser(userId int) error {
	args := mock.Called()

	return args.Error(0)
}

func (mock *MockUsersService) UpdateUser(userId int, userData *models.User) (user *models.User, err error) {
	args := mock.Called()

	result := args.Get(0)

	return result.(*models.User), args.Error(1)
}

func TestGetAllUsers(t *testing.T) {
	request, err := http.NewRequest("GET", "/users/", nil)
	response := httptest.NewRecorder()
	users := []*models.User{
		{
			ID:   1,
			Name: "TEST1",
		},
		{
			ID:   2,
			Name: "TEST2",
		},
	}

	mockService := new(MockUsersService)
	mockService.On("GetAllUsers").Return(users, nil)

	h := UserHandler{userService: mockService}
	h.GetAllUsers(response, request)

	assert.Equal(t, 200, response.Code)
	assert.NilError(t, err)
}

func TestCreateUser(t *testing.T) {
	user := models.User{
		Name:   "TEST1",
		Age:    100,
		Status: "TEST1",
	}

	userRes := models.User{
		ID:        1,
		Name:      "TEST1",
		Age:       100,
		Status:    "TEST1",
		CreatedAt: time.Now(),
	}

	byUser, _ := json.Marshal(&user)

	request, err := http.NewRequest("POST", "/users/", bytes.NewBuffer(byUser))
	response := httptest.NewRecorder()

	var userData *models.User = &userRes
	mockService := new(MockUsersService)
	mockService.On("AddUser").Return(userData, nil)

	h := UserHandler{userService: mockService}
	h.CreateUser(response, request)

	assert.Equal(t, 200, response.Code)
	assert.NilError(t, err)
}

func TestGetUser(t *testing.T) {
	request, err := http.NewRequest("GET", "/users/{userID}", nil)
	response := httptest.NewRecorder()
	user := models.User{
		ID:        1,
		Name:      "TEST1",
		Age:       100,
		Status:    "TEST1",
		CreatedAt: time.Now(),
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	var userData *models.User = &user
	mockService := new(MockUsersService)
	mockService.On("GetUser").Return(userData, nil)

	h := UserHandler{userService: mockService}
	h.GetUser(response, request)
	// fmt.Println(request)
	// fmt.Println(response)
	assert.Equal(t, 200, response.Code)
	assert.NilError(t, err)
}

func TestUpdateUser(t *testing.T) {
	user := models.User{
		ID:        1,
		Name:      "TEST1",
		Age:       100,
		Status:    "TEST1",
		CreatedAt: time.Now(),
	}

	byUser, _ := json.Marshal(&user)

	request, err := http.NewRequest("POST", "/users/{userID}", bytes.NewBuffer(byUser))
	response := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	var userData *models.User = &user
	mockService := new(MockUsersService)
	mockService.On("UpdateUser").Return(userData, nil)

	h := UserHandler{userService: mockService}
	h.UpdateUser(response, request)

	assert.Equal(t, 200, response.Code)
	assert.NilError(t, err)
}

func TestDeleteUser(t *testing.T) {
	request, err := http.NewRequest("DELETE", "/users/{userID}", nil)
	response := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	mockService := new(MockUsersService)
	mockService.On("DeleteUser").Return(nil)

	h := UserHandler{userService: mockService}
	h.DeleteUser(response, request)

	assert.Equal(t, 200, response.Code)
	assert.NilError(t, err)
}
