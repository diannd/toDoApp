package data

import (
	"toDoApp/features/note"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserID uint
	Title  string
	Note   string
	User   User `gorm:"foreignKey:UserID"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Note     []Note
}

func fromCore(dataCore note.Core) Note {
	dataModel := Note{
		UserID: dataCore.UserID,
		Title:  dataCore.Title,
		Note:   dataCore.Note,
	}
	return dataModel
}

func (data *Note) toCore() note.Core {
	return note.Core{
		ID:     data.ID,
		UserID: data.UserID,
		Title:  data.Title,
		Note:   data.Note,
	}
}

func toCoreList(data []Note) []note.Core {
	var dataCore []note.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
