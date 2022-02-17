package service

import (
	"gochicoba/models"
	"gochicoba/repository"
)

type UserService interface {
	GetAllUsers(models.UserFilter) ([]*models.User, error)
	GetUser(int) (*models.User, error)
	AddUser(*models.User) (*models.User, error)
	DeleteUser(int) error
	UpdateUser(int, *models.User) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) GetAllUsers(uf models.UserFilter) ([]*models.User, error) {
	list, err := us.userRepo.GetAllUsers(uf)
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

func (us *userService) AddUser(userData *models.User) (*models.User, error) {

	user, err := us.userRepo.AddUser(userData)

	if err != nil {
		return nil, err
	}

	return user, nil
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
