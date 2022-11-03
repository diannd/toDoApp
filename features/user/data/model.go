package data

import (
	"toDoApp/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Note     []Note
}

type Note struct {
	gorm.Model
	UserID uint
	Title  string
	Note   string
}

func fromCore(dataCore user.Core) User {
	dataModel := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
	}
	return dataModel
}

func (data *User) toCore() user.Core {
	return user.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}
