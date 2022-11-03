package data

import (
	"time"
	"toDoApp/features/agenda"

	"gorm.io/gorm"
)

type Agenda struct {
	gorm.Model
	UserID   uint
	Title    string
	Time     time.Time
	Desc     string
	Priority string
	User     User `gorm:"foreignKey:UserID"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Agenda   []Agenda
}

func fromCore(dataCore agenda.Core) Agenda {
	dataModel := Agenda{
		UserID:   dataCore.UserID,
		Title:    dataCore.Title,
		Time:     dataCore.Time,
		Desc:     dataCore.Desc,
		Priority: dataCore.Priority,
	}
	return dataModel
}

func (data *Agenda) toCore() agenda.Core {
	return agenda.Core{
		ID:       data.ID,
		UserID:   data.UserID,
		Title:    data.Title,
		Time:     data.Time,
		Desc:     data.Desc,
		Priority: data.Priority,
	}
}

func toCoreList(data []Agenda) []agenda.Core {
	var dataCore []agenda.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
