package repository 

import (
	"log"
	"belajar-api/app/features/book"
	"gorm.io/gorm"
)

type BookModel struct {
	db *gorm.DB
}

func New(d *gorm.DB) book.Repository {
	return &BookModel{
		db: d,
	}
}

func (bm *BookModel) Insert(newBook book.Core, userID string) (book.Core, error) {
		var insertData Book 
		insertData.Judul = newBook.Judul 
		insertData.Penerbit = newBook.Penerbit
		insertData.Tahun = newBook.Tahun
		insertData.UserID = userID

		if err := bm.db.Table("books").Create(&insertData).Error; err != nil {
			log.Println("Terjadi error saat create Book", err.Error())
			return book.Core{}, err
		}

		return newBook, nil 
}

func (bm *BookModel) GetAll() (any, error) {
	type ExpectedRespond struct {
		Judul string `json:"judul"`
		Tahun string `json:"tahun"`
		Nama string `json:"nama"`
	}
	res := []ExpectedRespond{}

	err := bm.db.Table("books").Select("books.judul as judul, books.tahun, users.nama as nama").Joins("JOIN users on users.hp = books.user_id").Scan(&res).Error

	if err != nil {
		log.Println("Terjadi error saat select Book", err.Error())
		return nil, err
	}
	return res, nil
}

func (bm *BookModel) GetBookByID(id uint) (Book, error) {
	res := Book{}

	if err := bm.db.Find(&res, id).Error; err != nil {
		log.Println("Terjadi error saat select Book", err.Error())
		return Book{}, err
	}
	return res, nil 
}