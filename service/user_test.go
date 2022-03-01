package service

import (
	"errors"
	"gochicoba/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (urm *UserRepoMock) GetAllUsers(uf models.UserFilter) ([]*models.User, error) {
	arg := urm.Called(uf)
	if arg.Get(0) == nil {
		return arg.Get(0).([]*models.User), nil
	}
	return arg.Get(0).([]*models.User), nil
}

func (urm *UserRepoMock) GetUser(iu int) (*models.User, error) {
	arg := urm.Called(iu)
	if arg.Get(0) == nil {
		return nil, errors.New("wrong id user")
	}
	usr := arg.Get(0).(*models.User)
	return usr, nil
}

func (urm *UserRepoMock) AddUser(mu *models.User) (*models.User, error) {
	arg := urm.Called(mu)
	if arg.Get(0) == nil {
		return nil, errors.New("wrong id user")
	}
	return arg.Get(1).(*models.User), nil
}

func (urm *UserRepoMock) DeleteUser(iu int) error {
	arg := urm.Called(iu)
	if arg.Get(0) == nil {
		return errors.New("wrong id user")
	}
	return arg.Error(0)
}

func (urm *UserRepoMock) UpdateUser(iu int, mu *models.User) (*models.User, error) {
	arg := urm.Called(iu, mu)
	if arg.Get(0) == nil {
		return nil, errors.New("wrong id user")
	}
	return arg.Get(1).(*models.User), nil
}

// GetAllUsers(models.UserFilter) ([]*models.User, error) OK
// GetUser(int) (*models.User, error) OK
// AddUser(*models.User) (*models.User, error) OK
// DeleteUser(int) error OK
// UpdateUser(int, *models.User) (*models.User, error) OK

var sliceUser = []*models.User{{
	ID:        1,
	Name:      "First",
	Age:       24,
	Status:    "First",
	CreatedAt: time.Now(),
}, {
	ID:        2,
	Name:      "Second",
	Age:       25,
	Status:    "Second",
	CreatedAt: time.Now(),
}}

var user = models.User{
	ID:        1,
	Name:      "First",
	Age:       24,
	Status:    "First",
	CreatedAt: time.Now(),
}

var userFil = models.UserFilter{
	Name:    "",
	AgeUp:   0,
	AgeDown: 30,
	Status:  "",
}

var userRepo = &UserRepoMock{Mock: mock.Mock{}}
var userServ = userService{userRepo: userRepo}

// func TestUserService_AddUser(t *testing.T) {
// 	userRepo.Mock.On("")
// }

func TestUserService_GetAllUsers(t *testing.T) {
	userRepo.Mock.On("GetAllUsers", userFil).Return(sliceUser, nil)
	usr, err := userServ.GetAllUsers(userFil)

	assert.Nil(t, err)
	assert.NotNil(t, usr)
}

func TestUserService_GetUser(t *testing.T) {
	userRepo.Mock.On("GetUser", 1).Return(&user, nil)
	usr, err := userServ.GetUser(1)

	assert.Nil(t, err)
	assert.NotNil(t, usr)
}

func TestUserService_GetUserNotFound(t *testing.T) {
	er := errors.New("user not found")
	userRepo.Mock.On("GetUser", 2).Return(nil, er)
	user, err := userServ.GetUser(2)

	assert.Nil(t, user)
	assert.NotNil(t, err)
}
