package config 

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

var (
	SECRET_JWT string = ""
)

func InitSQL() *gorm.DB {
	connStr:= fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			"root",
		    "golang",
	        "localhost",
            3306,
            "orm")
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect database, ", err.Error())
	}
	return db
}