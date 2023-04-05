package handler

import (
	"belajar-api/app/features/book"
	"belajar-api/helper"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	srv book.UseCase
}


func New(s book.UseCase) book.Handler{
	return &BookController{
		srv: s,
    }
}

func (bc *BookController) AddBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := helper.DecodeJWT(c.Get("user").(*jwt.Token))

		Input := &BookRequest{}
			if err := c.Bind(&Input); err != nil {
				c.Logger().Error("terjadi kesalahan bind", err.Error())
				return c.JSON(helper.ReponsFormat(http.StatusBadRequest,"terdapat kesalahan input dari book", nil))
			}

			res, err := bc.srv.AddBook(book.Core{Penerbit: Input.Judul, Tahun: Input.Tahun}, userID)

			if err != nil {
				c.Logger().Error("terjadi kesalahan", err.Error())
				return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server ", nil))
			}

			return c.JSON(helper.ReponsFormat(http.StatusCreated,"sukses menambahkan data", res))
	}
}

func (bc *BookController)GetAllBook() echo.HandlerFunc {
	return func (c echo.Context) error {
		res, err := bc.srv.GetAllBook()

		if err != nil {
			c.Logger().Error("Book model error", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan data", res))
	}
}