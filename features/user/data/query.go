package data

import (
	"toDoApp/features/user"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}
}

func (repo *userData) RegisterUser(data user.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *userData) ReadUser(id int) (user.Core, error) {
	var data User

	tx := repo.db.First(&data, id)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	return data.toCore(), nil
}

func (repo *userData) UpdateUser(id int, data user.Core) (int, error) {
	dataModel := fromCore(data)

	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}

func (repo *userData) DeleteData(id int) (int, error) {
	var data User

	tx := repo.db.Where("id = ?", id).Delete(&data)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}
