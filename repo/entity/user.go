package entity

import "github.com/jinzhu/gorm"

// User store the data related to the user.
type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index"`
	Name     string `gorm:"type:varchar(100)"`
	Password string
}
