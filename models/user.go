package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/skantuz/jwt-go"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	exp    time.Time
	jwt.StandardClaims
}

// User a struct to rep user us
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique" `
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique"`
	Name     string `json:"name" json:"name"`
	Lastname string `json:"lastname"`
	Token    string `json:"token"`
	RoleId   uint   `json:"role_id" gorm:"index:role_id"`
	Role     Role   `json:"role" gorm:"foreignkey:id;association_foreignkey:role_id"`
}
