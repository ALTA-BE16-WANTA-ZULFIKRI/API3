package usecase_test

import (
	"belajar-api/app/features/user"
	"belajar-api/app/features/user/mocks"
	"belajar-api/app/features/user/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	repo := mocks.NewRepository(t)
	uc := usecase.New(repo)
	succesCaseData := user.Core{Nama: "jerry", HP: "12345", Password: "tonohaha577"}

	t.Run("Sukses login", func(t *testing.T) {

		repo.On("Login", succesCaseData.HP, succesCaseData.Password).Return(user.Core{Nama: "jerrry", HP: "12345"}, nil).Once()

		result, err := uc.Login("12345", "tonohaha577")

		assert.Nil(t, err)
		assert.Equal(t, "12345", result.HP)
		repo.AssertExpectations(t)
	})

	t.Run("Password salah", func(t *testing.T) {
		repo.On("Login", succesCaseData.HP, "bangsat").Return(user.Core{}, errors.New("Password Salah")).Once()

		result, err := uc.Login(succesCaseData.HP, "bangsat")

		assert.Error(t, err)
		assert.Empty(t, "", result.Nama)
		repo.AssertExpectations(t)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		repo.On("Login", "6789", "tonohaha").Return(user.Core{}, errors.New("data tidak ditemukan")).Once()
		result, err := uc.Login("6789", "tonohaha")

		assert.Error(t, err)
		assert.ErrorContainsf(t, err, "data tidak ditemukan", "errir data tdaj dutenykan")
		assert.Empty(t, result.Nama)
		repo.AssertExpectations(t)
	})

	t.Run("kesalahan pada server", func(t *testing.T) {
		repo.On("Login", succesCaseData.HP, "tonohaha").Return(user.Core{}, errors.New("column not exist")).Once()

		result, err := uc.Login("12345", "tonohaha")

		assert.Error(t, err)
		assert.ErrorContainsf(t, err, result.Nama, "Keasalahan pada Server")
		repo.AssertExpectations(t)
	})

}

func TestRegister(t *testing.T) {
	repo := mocks.NewRepository(t)
	uc := usecase.New(repo)
	succesCaseData := user.Core{Nama: "jerry", HP: "12345", Password: "tonohaha577"}

	t.Run("sukses Register", func(t *testing.T) {
		repo.On("Insert", succesCaseData).Return(succesCaseData, nil).Once()
		err := uc.Register(succesCaseData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed when func insert return error", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(user.Core{}, errors.New("error insert data")).Once()
		err := uc.Register(user.Core{})
		assert.NotNil(t, err)
		assert.Equal(t, "error insert data", err.Error())
	})

	t.Run("Failel validate", func(t *testing.T) {
		inpuData := user.Core{
			Nama:     "joko",
			HP:       "6666",
			Password: "jok",
		}
		repo.On("Insert", mock.Anything).Return(user.Core{}, errors.New("failed validate")).Once()
		err := uc.Register(inpuData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)

	})
}
