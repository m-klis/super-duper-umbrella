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
	args := urm.Called(mu)
	if args.Get(0) == nil {
		return nil, errors.New("wrong id user")
	}
	usr := args.Get(0).(*models.User)
	return usr, nil
}

func (urm *UserRepoMock) DeleteUser(iu int) error {
	arg := urm.Called()
	return arg.Error(0)
}

func (urm *UserRepoMock) UpdateUser(iu int, mu *models.User) (*models.User, error) {
	arg := urm.Called(iu, mu)
	if arg.Get(0) == nil {
		return nil, errors.New("wrong id user")
	}
	return arg.Get(0).(*models.User), arg.Error(1)
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

func TestUserService_GetAllUsers(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	userRepo.Mock.On("GetAllUsers", userFil).Return(sliceUser, nil)
	usr, err := userServ.GetAllUsers(userFil)

	assert.Nil(t, err)
	assert.NotNil(t, usr)
}

func TestUserService_GetUser(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	userRepo.Mock.On("GetUser", 1).Return(&user, nil)
	usr, err := userServ.GetUser(1)

	assert.Nil(t, err)
	assert.NotNil(t, usr)
}

func TestUserService_GetUserNotFound(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	er := errors.New("user not found")
	userRepo.Mock.On("GetUser", 2).Return(nil, er)
	userr, err := userServ.GetUser(2)

	assert.Nil(t, userr)
	assert.NotNil(t, err)
}

func TestUserService_AddUser(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	userRepo.Mock.On("AddUser", &user).Return(&user, nil)

	userr, err := userServ.AddUser(&user)

	assert.Nil(t, err, "Error must be nill")
	assert.Equal(t, &user, userr)
}

func TestUserService_AddUserFail(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	newuser := models.User{
		ID:        1,
		Name:      "First",
		Age:       24,
		Status:    "First",
		CreatedAt: time.Now(),
	}
	erro := errors.New("data exits")
	userRepo.Mock.On("AddUser", &newuser).Return(nil, erro)

	userr, err := userServ.AddUser(&newuser)

	assert.NotNil(t, err, "must error")
	assert.Nil(t, userr, "must nil")
}

func TestUserService_DeleteUser(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	newuser := models.User{
		ID:        1,
		Name:      "First",
		Age:       24,
		Status:    "First",
		CreatedAt: time.Now(),
	}

	userRepo.Mock.On("DeleteUser").Return(nil)

	err := userServ.DeleteUser(newuser.ID)

	assert.Nil(t, err, "err must be nil")
}

func TestUserService_DeleteUserFail(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	newuser := models.User{
		ID:        1,
		Name:      "First",
		Age:       24,
		Status:    "First",
		CreatedAt: time.Now(),
	}

	userRepo.Mock.On("DeleteUser").Return(errors.New("wrong id"))

	err := userServ.DeleteUser(newuser.ID)

	assert.Error(t, err, "must be error")
}

func TestUserService_UpdateUser(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	newuser := models.User{
		ID:        1,
		Name:      "First",
		Age:       24,
		Status:    "First",
		CreatedAt: time.Now(),
	}

	newuser1 := models.User{
		ID:        1,
		Name:      "First",
		Age:       24,
		Status:    "First",
		CreatedAt: time.Now(),
	}

	userRepo.Mock.On("UpdateUser", newuser.ID, &newuser1).Return(&newuser, nil)

	res, err := userServ.UpdateUser(newuser.ID, &newuser1)

	assert.Nil(t, err, "error must be nil")
	assert.Equal(t, newuser.Name, res.Name)

}