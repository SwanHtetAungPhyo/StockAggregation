
package models

import (
"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string           `gorm:"type:varchar(100);" json:"name"`
	Email           string           `gorm:"type:varchar(100);unique" json:"email"`
	StockWatchLists []StockWatchList `gorm:"foreignKey:UserID" json:"stockWatchLists"`
}


type StockWatchList struct {
	gorm.Model
	UserID   uint   `json:"userId"`
	Stock    string `gorm:"type:varchar(100);" json:"stock"`
	Quantity int    `json:"quantity"`
}