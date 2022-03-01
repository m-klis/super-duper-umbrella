package service

import (
	"gochicoba/models"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	Mock mock.Mock
}

func (urm *UserRepoMock) GetAllUsers(uf models.UserFilter) ([]*models.User, error) {
	// arguments := urm.Mock.Called(uf)
	return nil, nil
}

func (urm *UserRepoMock) GetUser(iu int) (*models.User, error) {

	return nil, nil
}

func (urm *UserRepoMock) AddUser(mu *models.User) (*models.User, error) {

	return nil, nil
}

func (urm *UserRepoMock) DeleteUser(iu int) error {

	return nil
}

func (urm *UserRepoMock) UpdateUser(iu int, mu *models.User) (*models.User, error) {

	return nil, nil
}

// GetAllUsers(models.UserFilter) ([]*models.User, error) OK
// GetUser(int) (*models.User, error) OK
// AddUser(*models.User) (*models.User, error) OK
// DeleteUser(int) error OK
// UpdateUser(int, *models.User) (*models.User, error) OK
