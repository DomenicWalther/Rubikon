package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id" gorm:"not null"`
	ImageURL    string `json:"image_url"`
	Users       []User `gorm:"many2many:group_users;"`
	Is_Public   bool   `json:"is_public" gorm:"not null;default:true"`
}
