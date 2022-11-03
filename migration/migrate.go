package migration

import (
	noteModel "toDoApp/features/note/data"
	userModel "toDoApp/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&noteModel.Note{})
}
