package usecase

import (
	"errors"
	"belajar-api/app/features/book"
	"strings"

	"github.com/labstack/gommon/log"
)

type BookModel struct {
	repo book.Repository
}

func New(br book.Repository) book.UseCase {
	return &BookModel{
		repo: br,
	}
}

func (bm *BookModel) AddBook(newBook book.Core, user_id string) (book.Core, error) {
		result, err := bm.repo.Insert(newBook, user_id)
		if err != nil {
			log.Error("terjadi kesalahan input buku", err.Error())
			if strings.Contains(err.Error(), "too much") {
				return book.Core{}, errors.New("terdapat kesalahan input, nilai yang diberikan terlalu panjang")
			}
			return book.Core{}, errors.New("terdapat masalah pada server")
		}
		return result,nil
}

func (bm *BookModel) GetAllBook() (any, error) {
	result, err := bm.repo.GetAll()
	if err != nil {
		log.Error("terjadi kesalahan get buku", err.Error())
		return book.Core{}, errors.New("terdapat masalah pada server")
	}
	return result,nil
}