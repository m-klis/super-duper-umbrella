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

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{userService: userService}
}

func (ih *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	list, err := ih.userService.GetAllUsers()
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", list)
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", list)
	return
}

func (ih *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "UserID")

	UserIDInt, err := strconv.Atoi(UserID)
	fmt.Println(UserID)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusBadRequest, "id must be integer", err)
		return
	}
	User, err := ih.userService.GetUser(UserIDInt)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}
	if User == nil {
		helpers.ErrorResponse(w, r, http.StatusNotFound, "not found", err)
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", User)
	return
}

func (ih *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var User *models.User
	err := json.NewDecoder(r.Body).Decode(&User)
	fmt.Println(User)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}

	UserId, err := ih.userService.AddUser(User)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", UserId)
	return
}

func (ih *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "userID")
	UserIDInt, err := strconv.Atoi(UserID)
	//fmt.Println(UserID)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusBadRequest, "id must be integer", err)
		return
	}
	var User *models.User
	err = json.NewDecoder(r.Body).Decode(&User)
	//fmt.Println(User)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}

	User, err = ih.userService.UpdateUser(UserIDInt, User)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}
	helpers.CustomResponse(w, r, http.StatusOK, "success", User)
	return
}

func (ih *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "UserID")
	UserIDInt, err := strconv.Atoi(UserID)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}

	err = ih.userService.DeleteUser(UserIDInt)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err)
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", nil)
	return
}
