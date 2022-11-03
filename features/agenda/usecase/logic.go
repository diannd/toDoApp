package usecase

import (
	"errors"
	"toDoApp/features/agenda"
)

type agendaUsecase struct {
	agendaData agenda.DataInterface
}

func New(data agenda.DataInterface) agenda.UsecaseInterface {
	return &agendaUsecase{
		agendaData: data,
	}
}

func (usecase *agendaUsecase) PostAgenda(data agenda.Core) (int, error) {
	if data.Title == "" {
		return -1, errors.New("data tidak boleh kosong")
	}

	row, err := usecase.agendaData.CreateAgenda(data)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *agendaUsecase) GetAgenda(id int) ([]agenda.Core, error) {
	data, err := usecase.agendaData.ReadAgenda(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (usecase *agendaUsecase) GetAgendaByID(id int) (agenda.Core, error) {
	data, err := usecase.agendaData.ReadAgendaByID(id)
	if err != nil {
		return agenda.Core{}, err
	}

	return data, nil
}

func (usecase *agendaUsecase) PutAgenda(id int, data agenda.Core) (int, error) {
	row, err := usecase.agendaData.UpdateAgenda(id, data)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *agendaUsecase) DeleteAgenda(id int) (int, error) {
	row, err := usecase.agendaData.DeleteData(id)
	if err != nil {
		return -1, err
	}

	return row, nil
}
