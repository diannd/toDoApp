package data

import (
	"toDoApp/features/agenda"

	"gorm.io/gorm"
)

type agendaData struct {
	db *gorm.DB
}

func New(db *gorm.DB) agenda.DataInterface {
	return &agendaData{
		db: db,
	}
}

func (repo *agendaData) CreateAgenda(data agenda.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *agendaData) ReadAgenda(token int) ([]agenda.Core, error) {
	var data []Agenda
	tx := repo.db.Model(&Agenda{}).Where("user_id = ?", token).Preload("User").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return toCoreList(data), nil
}

func (repo *agendaData) ReadAgendaByID(id int) (agenda.Core, error) {
	var data Agenda
	tx := repo.db.Model(&Agenda{}).Where("id = ?", id).Preload("User").Find(&data)
	if tx.Error != nil {
		return agenda.Core{}, tx.Error
	}
	return data.toCore(), nil
}

func (repo *agendaData) UpdateAgenda(id int, data agenda.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Model(&Agenda{}).Where("id = ?", id).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}

func (repo *agendaData) DeleteData(id int) (int, error) {
	var deleteData Agenda

	tx := repo.db.Where("id = ?", id).Delete(&deleteData)
	if tx.Error != nil {
		return -1, tx.Error
	}

	return 1, nil
}
