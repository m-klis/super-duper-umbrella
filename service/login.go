package service

import (
	"gochicoba/models"
	"gochicoba/repository"
)

type LoginService interface {
	CheckLogin(models.Credentials) error
}

type loginService struct {
	loginRepo repository.LoginRepository
}

func NewLoginService(loginRepo repository.LoginRepository) LoginService {
	return &loginService{
		loginRepo: loginRepo,
	}
}

func (ls *loginService) CheckLogin(mc models.Credentials) error {
	err := ls.loginRepo.CheckLogin(mc)
	if err != nil {
		return err
	}
	return nil
}
