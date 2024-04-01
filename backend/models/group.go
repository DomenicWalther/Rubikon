package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id" gorm:"not null"`
	ImageURL    string `json:"imageURL"`
	Users       []User `gorm:"many2many:group_users;"`
	Is_Public   bool   `json:"isPrivate" gorm:"not null;default:true"`
}
