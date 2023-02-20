package model

import (
	"gorm.io/gorm"
)

type UrlModel struct {
	Url string `json:url`
}
type TinyUrlData struct {
	gorm.Model
	Hash string
	Url  string
}

// gorm.Model equals
// type User struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// 	Name string
//   }
