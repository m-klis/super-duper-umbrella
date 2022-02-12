package service

import (
	"gochicoba/models"
	"gochicoba/repository"
)

type UserService interface {
	GetAllUsers() ([]*models.User, error)
	GetUser(userId int) (userData *models.User, err error)
	AddUser(user *models.User) (userId int, err error)
	DeleteUser(userId int) error
	UpdateUser(userId int, userData *models.User) (user *models.User, err error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) GetAllUsers() ([]*models.User, error) {
	list, err := us.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (us *userService) GetUser(id int) (*models.User, error) {
	item, err := us.userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (us *userService) AddUser(userData *models.User) (int, error) {

	idUser, err := us.userRepo.AddUser(userData)

	if err != nil {
		return 0, err
	}

	return idUser, nil
}

func (us *userService) UpdateUser(userId int, userData *models.User) (*models.User, error) {

	user, err := us.userRepo.UpdateUser(userId, userData)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) DeleteUser(userId int) error {
	err := us.userRepo.DeleteUser(userId)
	if err != nil {
		return err
	}

	return nil
}
