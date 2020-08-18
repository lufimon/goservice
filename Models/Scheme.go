package Models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID        uint    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"userId"`
	FirstName string  `gorm:"column:first_name" json:"firstName"`
	LastName  string  `gorm:"column:last_name" json:"lastName"`
	BirthDay  string  `gorm:"column:birth_day" json:"birthDay"`
	Age       int64   `gorm:"column:age" json:"age"`
	Contact   Contact `gorm:"-" json:"contact"`
	Facebook  string  `gorm:"column:facebook"`
	LineId    string  `gorm:"column:line_id"`
	Instagram string  `gorm:"column:instagram"`
}

type Users struct {
	ID        uint    `json:"userId"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	BirthDay  string  `json:"birthDay"`
	Age       int64   `json:"age"`
	Contact   Contact `json:"contact"`
}

type Contact struct {
	Facebook  string `json:"facebook"`
	LineId    string `json:"lineId"`
	Instagram string `json:"instagram"`
}

func (u *User) TableName() string {
	return "user"
}
