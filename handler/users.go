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
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{userService: userService}
}

func (ih *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	ageup := r.URL.Query().Get("ageup")
	agedown := r.URL.Query().Get("agedown")
	status := r.URL.Query().Get("status")

	if ageup == "" {
		ageup = "0"
	}
	au, err := strconv.Atoi(ageup)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusBadRequest, "id must be integer", err.Error())
		return
	}

	if agedown == "" {
		agedown = "0"
	}
	ad, err := strconv.Atoi(agedown)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusBadRequest, "id must be integer", err.Error())
		return
	}

	filter := models.UserFilter{
		Name:    name,
		AgeUp:   au,
		AgeDown: ad,
		Status:  status,
	}

	fmt.Println(filter.Name, " + ", filter.Status)

	list, err := ih.userService.GetAllUsers(filter)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", list)
	return
}

func (ih *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "userID")
	UserIDInt, err := strconv.Atoi(UserID)
	//fmt.Println(UserID)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusBadRequest, "id must be integer", err.Error())
		return
	}
	User, err := ih.userService.GetUser(UserIDInt)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	if User == nil {
		helpers.ErrorResponse(w, r, http.StatusNotFound, "not found", err.Error())
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", User)
	return
}

func (ih *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	// fmt.Println(User)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	ud, err := ih.userService.AddUser(user)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	response := models.UserResponse{
		ID:        ud.ID,
		Name:      ud.Name,
		Age:       ud.Age,
		Status:    ud.Status,
		CreatedAt: helpers.ConvertMonth(ud.CreatedAt),
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", response)
	return
}

func (ih *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "userID")
	UserIDInt, err := strconv.Atoi(UserID)
	//fmt.Println(UserID)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusBadRequest, "id must be integer", err.Error())
		return
	}
	var User *models.User
	err = json.NewDecoder(r.Body).Decode(&User)
	//fmt.Println(User)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	User, err = ih.userService.UpdateUser(UserIDInt, User)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	helpers.CustomResponse(w, r, http.StatusOK, "success", User)
	return
}

func (ih *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "userID")
	UserIDInt, err := strconv.Atoi(UserID)

	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	err = ih.userService.DeleteUser(UserIDInt)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	helpers.CustomResponse(w, r, http.StatusOK, "success", nil)
	return
}
