package usecase

import (
	"errors"
	"toDoApp/features/user"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) PostUser(data user.Core) (int, error) {
	if data.Name == "" || data.Email == "" || data.Password == "" {
		return -1, errors.New("data tidak boleh kosong")
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	data.Password = string(hashPass)
	row, err := usecase.userData.RegisterUser(data)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *userUsecase) GetUser(id int) (user.Core, error) {
	data, err := usecase.userData.ReadUser(id)
	if err != nil {
		return user.Core{}, err
	}

	return data, nil
}

func (usecase *userUsecase) PutUser(id int, newData user.Core) (int, error) {

	if newData.Password != "" {
		hashPass, err := bcrypt.GenerateFromPassword([]byte(newData.Password), bcrypt.DefaultCost)
		if err != nil {
			return -1, err
		}
		newData.Password = string(hashPass)
	}

	row, err := usecase.userData.UpdateUser(id, newData)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (usecase *userUsecase) DeleteUser(id int) (int, error) {
	row, err := usecase.userData.DeleteData(id)
	if err != nil {
		return -1, err
	}
	return row, nil
}
