package models

import (
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

type User struct {
	gorm.Model
	Name            string           `gorm:"type:varchar(100);" json:"name" validate:"required,min=3"`
	Email           string           `gorm:"type:varchar(100);unique" json:"email" validate:"required,email"`
	Password        string           `gorm:"type:varchar(100);" json:"password" validate:"required,min=6"`
	StockWatchLists []StockWatchList `gorm:"foreignKey:UserID" json:"stockWatchLists,omitempty"`
}

type StockWatchList struct {
	gorm.Model
	UserID   uint   `json:"userId"`
	Stock    string `gorm:"type:varchar(100);" json:"stock" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,min=1"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *User) Validate() error {
	return validate.Struct(u)
}

func (s *StockWatchList) Validate() error {
	return validate.Struct(s)
}
