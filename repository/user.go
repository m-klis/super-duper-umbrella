package repository

import (
	"errors"
	"gochicoba/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(models.UserFilter) ([]*models.User, error)
	GetUser(int) (*models.User, error)
	AddUser(*models.User) (*models.User, error)
	DeleteUser(int) error
	UpdateUser(int, *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetAllUsers(uf models.UserFilter) (userList []*models.User, err error) {
	var list []*models.User
	query := ur.db.Debug()

	if uf.Name != "" {
		query = query.Where("name LIKE ?", "%"+uf.Name+"%")
	}

	if uf.Status != "" {
		query = query.Where("status LIKE ?", "%"+uf.Status+"%")
	}

	if uf.AgeUp != 0 && uf.AgeDown != 0 {
		query = query.Where("age BETWEEN ? AND ?", uf.AgeDown, uf.AgeUp)
	}

	err = query.Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (ur *userRepository) GetUser(id int) (user *models.User, err error) {
	query := ur.db
	err = query.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) AddUser(userData *models.User) (*models.User, error) {
	query := ur.db
	err := query.Create(&userData).Error
	//fmt.Println(item)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (ur *userRepository) UpdateUser(userId int, userData *models.User) (user *models.User, err error) {
	query := ur.db
	//field := &itemData
	err = query.Model(&userData).Where("id", userId).Updates(&userData).Error
	if err != nil {
		return nil, err
	}
	err = query.Where("id = ?", userId).First(&userData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return userData, nil
}

func (ur *userRepository) DeleteUser(userId int) error {
	query := ur.db
	var user *models.User
	err := query.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return err
	}
	err = query.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
