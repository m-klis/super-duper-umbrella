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
		return nil, arg.Error(1)
	}
	return arg.Get(0).([]*models.User), nil
}

func (urm *UserRepoMock) GetUser(iu int) (*models.User, error) {
	arg := urm.Called(iu)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	}
	usr := arg.Get(0).(*models.User)
	return usr, nil
}

func (urm *UserRepoMock) AddUser(mu *models.User) (*models.User, error) {
	args := urm.Called(mu)
	if args.Get(0) == nil {
		return nil, args.Error(1)
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

func TestUserService_GetAllUsers(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	var userFil = models.UserFilter{
		Name:    "",
		AgeUp:   0,
		AgeDown: 30,
		Status:  "",
	}

	userRepo.Mock.On("GetAllUsers", userFil).Return(sliceUser, nil)
	usr, err := userServ.GetAllUsers(userFil)
	userRepo.Mock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.NotNil(t, usr)
}

func TestUserService_GetAllUsersFail(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	var userFil = models.UserFilter{
		Name:    "",
		AgeUp:   0,
		AgeDown: 30,
		Status:  "",
	}

	userRepo.Mock.On("GetAllUsers", userFil).Return(nil, errors.New("failed get data"))
	usr, err := userServ.GetAllUsers(userFil)
	userRepo.Mock.AssertExpectations(t)

	assert.Nil(t, usr)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestUserService_GetUser(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	userRepo.Mock.On("GetUser", 1).Return(&user, nil)
	usr, err := userServ.GetUser(1)
	userRepo.Mock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.NotNil(t, usr)
}

func TestUserService_GetUserNotFound(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	er := errors.New("user not found")
	userRepo.Mock.On("GetUser", 2).Return(nil, er)
	userr, err := userServ.GetUser(2)
	userRepo.Mock.AssertExpectations(t)

	assert.Nil(t, userr)
	assert.NotNil(t, err)
}

func TestUserService_AddUser(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	userRepo.Mock.On("AddUser", &user).Return(&user, nil)
	userr, err := userServ.AddUser(&user)
	userRepo.Mock.AssertExpectations(t)

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
	userRepo.Mock.AssertExpectations(t)

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
	userRepo.Mock.AssertExpectations(t)

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
	userRepo.Mock.AssertExpectations(t)

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
	userRepo.Mock.AssertExpectations(t)

	assert.Nil(t, err, "error must be nil")
	assert.Equal(t, newuser.Name, res.Name)

}

func TestUserService_UpdateUserFail(t *testing.T) {
	var userRepo = &UserRepoMock{Mock: mock.Mock{}}
	var userServ = userService{userRepo: userRepo}

	newuser1 := models.User{
		ID:        1,
		Name:      "First",
		Age:       24,
		Status:    "First",
		CreatedAt: time.Now(),
	}

	userRepo.Mock.On("UpdateUser", 2, &newuser1).Return(nil, errors.New("id not match"))
	res, err := userServ.UpdateUser(2, &newuser1)
	userRepo.Mock.AssertExpectations(t)

	assert.Error(t, err, "must return error")
	assert.Nil(t, res, "must return nil")

}
