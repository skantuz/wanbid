package models

import (
	"backend/app"

	"github.com/jinzhu/gorm"
)

//a struct to rep Role
type Role struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `gorm:"foreignkey:role_id"`
}

type Roles []Role

//Validate incoming user details...

func (role *Role) Create() map[string]interface{} {

	GetDB().Create(role)

	response := app.Message(true, "Role has been created")
	response["role"] = role
	return response
}

func GetRole(u uint) *Role {

	acc := &Role{}
	GetDB().Table("roles").Where("id = ?", u).First(acc)
	if acc.Name == "" { //Role not found!
		return nil
	}
	return acc
}

func ListRoles() *Roles {

	roles := &Roles{}
	err := GetDB().Find(roles).Error
	if err != nil { //Roles not found!
		return nil
	}

	return roles
}
