package factory

import (
	authData "toDoApp/features/auth/data"
	authDelivery "toDoApp/features/auth/delivery"
	authUsecase "toDoApp/features/auth/usecase"

	userData "toDoApp/features/user/data"
	userDelivery "toDoApp/features/user/delivery"
	userUsecase "toDoApp/features/user/usecase"

	noteData "toDoApp/features/note/data"
	noteDelivery "toDoApp/features/note/delivery"
	noteUsecase "toDoApp/features/note/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	authDataFactory := authData.New(db)
	authUsecaseFactory := authUsecase.New(authDataFactory)
	authDelivery.New(e, authUsecaseFactory)

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	noteDataFactory := noteData.New(db)
	noteUsecaseFactory := noteUsecase.New(noteDataFactory)
	noteDelivery.New(e, noteUsecaseFactory)
}
