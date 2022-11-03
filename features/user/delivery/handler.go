package delivery

import (
	"toDoApp/features/user"
	"toDoApp/middlewares"
	"toDoApp/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}

	e.POST("/users", handler.PostUser)
	e.GET("/users", handler.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", handler.PutUser, middlewares.JWTMiddleware())
	e.DELETE("/users", handler.DeleteUser, middlewares.JWTMiddleware())
}

func (delivery *UserDelivery) PostUser(c echo.Context) error {
	var dataRegister UserRequest

	errBind := c.Bind(&dataRegister)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error binda data"))
	}

	row, err := delivery.userUsecase.PostUser(toCore(dataRegister))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *UserDelivery) GetUser(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	res, err := delivery.userUsecase.GetUser(idToken)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error no data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("succes get data", res))
}

func (delivery *UserDelivery) PutUser(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	var updateData UserRequest

	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	row, err := delivery.userUsecase.PutUser(idToken, toCore(updateData))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error update data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success update data"))
}

func (delivery *UserDelivery) DeleteUser(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	row, err := delivery.userUsecase.DeleteUser(idToken)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error delete data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error delete data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success delete data"))

}
