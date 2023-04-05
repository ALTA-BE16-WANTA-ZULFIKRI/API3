package routes

import (
	"belajar-api/app/features/book"
	"belajar-api/app/features/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	
)

func Route(e *echo.Echo, uc user.Handler, bc book.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/login", uc.Login())
	e.POST("/users", uc.Register())

	e.GET("/books", bc.GetAllBook())


	e.POST("/books", bc.AddBook(), middleware.JWT([]byte("s3cr3t!!")))

	
}	