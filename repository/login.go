package repository

import (
	"errors"
	"gochicoba/models"

	"gorm.io/gorm"
)

type LoginRepository interface {
	CheckLogin(models.Credentials) error
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) LoginRepository {
	return &loginRepository{
		db: db,
	}
}

func (lr *loginRepository) CheckLogin(mc models.Credentials) error {
	var dataLogin models.Login
	query := lr.db
	err := query.Where("username = ?", mc.Username).First(&dataLogin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	if dataLogin.Username != mc.Username {
		return errors.New("wrong username")
	}

	if dataLogin.Password != mc.Password {
		return errors.New("wrong password")
	}

	return nil
}
