package migration

import (
	agendaModel "toDoApp/features/agenda/data"
	noteModel "toDoApp/features/note/data"
	userModel "toDoApp/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&noteModel.Note{})
	db.AutoMigrate(&agendaModel.Agenda{})
}
