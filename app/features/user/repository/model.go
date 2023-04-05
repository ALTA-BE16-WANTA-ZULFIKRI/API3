package repository

import (
	"belajar-api/app/features/book/repository"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID string 
	Nama string
	HP string `gorm:"primarykey;type:varchar(13);"`
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Books []repository.Book `gorm:"foreignKey:UserID"`
}