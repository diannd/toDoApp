package usecase

import (
	"errors"
	"toDoApp/features/note"
)

type noteUsecase struct {
	noteData note.DataInterface
}

func New(data note.DataInterface) note.UsecaseInterface {
	return &noteUsecase{
		noteData: data,
	}
}

func (usecase *noteUsecase) PostNote(data note.Core) (int, error) {
	if data.Title == "" || data.Note == "" {
		return -1, errors.New("data tidak boleh kosong")
	}

	row, err := usecase.noteData.CreateNote(data)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *noteUsecase) GetNoteByID(id int) (note.Core, error) {
	data, err := usecase.noteData.ReadNoteByID(id)
	if err != nil {
		return note.Core{}, err
	}

	return data, nil
}

func (usecase *noteUsecase) GetNote(id int) ([]note.Core, error) {
	data, err := usecase.noteData.ReadNote(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (usecase *noteUsecase) PutNote(id int, data note.Core) (int, error) {
	row, err := usecase.noteData.UpdateNote(id, data)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *noteUsecase) DeleteNote(id int) (int, error) {
	row, err := usecase.noteData.DeleteData(id)
	if err != nil {
		return -1, err
	}

	return row, nil
}
