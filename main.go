package main 

import (
	bhandle "belajar-api/app/features/book/handler"
	brepo "belajar-api/app/features/book/repository"
	blogic "belajar-api/app/features/book/usecase"
	uhandle "belajar-api/app/features/user/handler"
	urepo "belajar-api/app/features/user/repository"
	ulogic "belajar-api/app/features/user/usecase"

	"belajar-api/app/routes"
	"belajar-api/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitSQL()
	cfg.AutoMigrate(urepo.User{})
	cfg.AutoMigrate(brepo.Book{})

	mdl := urepo.New(cfg)
	srv := ulogic.New(mdl)
	ctl := uhandle.New(srv)

	bookMdl := brepo.New(cfg)
	bookSrv := blogic.New(bookMdl)
	bookCtl := bhandle.New(bookSrv)

	// ROUTING 
	routes.Route(e, ctl, bookCtl)

	e.Start(":8000")
}