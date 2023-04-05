package book 

import "github.com/labstack/echo/v4"

type Core struct {
	ID uint
	Judul string
	Tahun string
	Penerbit string
	Pemilik string
}

type Handler interface {
	AddBook() echo.HandlerFunc
	GetAllBook() echo.HandlerFunc
}

type UseCase interface {
	AddBook(newBook Core, user_id string) (Core, error)
	GetAllBook() (any, error)
}

type Repository interface {
	Insert(newBook Core, user_id string) (Core, error)
	GetAll() (any, error)
}