package data

import (
	"toDoApp/features/note"

	"gorm.io/gorm"
)

type noteData struct {
	db *gorm.DB
}

func New(db *gorm.DB) note.DataInterface {
	return &noteData{
		db: db,
	}
}

func (repo *noteData) CreateNote(data note.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *noteData) ReadNoteByID(id int) (note.Core, error) {
	var data Note
	tx := repo.db.Model(&Note{}).Where("id = ?", id).Preload("User").Find(&data)
	if tx.Error != nil {
		return note.Core{}, tx.Error
	}
	return data.toCore(), nil
}

func (repo *noteData) ReadNote(token int) ([]note.Core, error) {
	var data []Note
	tx := repo.db.Model(&Note{}).Where("user_id = ?", token).Preload("User").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return toCoreList(data), nil
}

func (repo *noteData) UpdateNote(id int, data note.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Model(&Note{}).Where("id = ?", id).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}

func (repo *noteData) DeleteData(id int) (int, error) {
	var deleteData Note

	tx := repo.db.Where("id = ?", id).Delete(&deleteData)
	if tx.Error != nil {
		return -1, tx.Error
	}

	return 1, nil
}
