package repository 

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Judul string `json:"judul"`
	Tahun string `json:"tahun"`
	Penerbit string `json:"penerbit"`
	UserID string `json:"hp" gorm:"type:varchar(13)"`
}