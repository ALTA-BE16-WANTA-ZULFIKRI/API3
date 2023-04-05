package usecase

import (
	"belajar-api/app/features/user"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
)

type UserLogic struct {
	m user.Repository
}

func New(r user.Repository) user.UseCase {
	return &UserLogic{
		m: r,
	}
}

func (ul *UserLogic) Login(hp string, password string) (user.Core, error) {
	result, err := ul.m.Login(hp, password)
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") {
			return user.Core{}, errors.New("data tidak ditemukan")
		} else if strings.Contains(err.Error(), "salah") {
			return user.Core{}, errors.New("password salah")
		}
		return user.Core{}, errors.New("terdapat permasalahan pada server")
	}

	return result, nil
}

func (ul *UserLogic) Register(newUser user.Core) error {
	_, err := ul.m.Insert(newUser)
	if err != nil {
		log.Error("register logic error:", err.Error())
		return err
	}

	return nil
}
