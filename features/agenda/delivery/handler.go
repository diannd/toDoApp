package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"toDoApp/features/agenda"
	"toDoApp/middlewares"
	"toDoApp/utils/helper"

	"github.com/labstack/echo/v4"
)

type AgendaDelivery struct {
	noteUsecase agenda.UsecaseInterface
}

func New(e *echo.Echo, usecase agenda.UsecaseInterface) {
	handler := &AgendaDelivery{
		noteUsecase: usecase,
	}

	e.POST("/agenda", handler.PostAgenda, middlewares.JWTMiddleware())
	e.GET("/agenda", handler.GetAgenda, middlewares.JWTMiddleware())
	e.GET("/agenda/:id", handler.GetAgendaByID, middlewares.JWTMiddleware())
	e.PUT("/agenda/:id", handler.PutAgenda, middlewares.JWTMiddleware())
	e.DELETE("/agenda/:id", handler.DeleteAgenda, middlewares.JWTMiddleware())
}

func (delivery *AgendaDelivery) PostAgenda(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	var dataAgenda AgendaRequest

	errBind := c.Bind(&dataAgenda)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	layout_date := "2006-01-02T15:04"
	timeAgenda, err := time.Parse(layout_date, fmt.Sprintf("%sT00:00", dataAgenda.Time))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed format time"))
	}

	if dataAgenda.Priority == "urgent" {
		dataAgenda.Priority = "urgent"
	} else if dataAgenda.Priority == "important" {
		dataAgenda.Priority = "important"
	} else if dataAgenda.Priority == "not urgent" {
		dataAgenda.Priority = "not urgent"
	} else if dataAgenda.Priority == "not important" {
		dataAgenda.Priority = "not important"
	} else {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("priority does not match"))
	}

	dataCore := toCore(dataAgenda)
	dataCore.Time = timeAgenda
	dataCore.UserID = uint(idToken)

	row, err := delivery.noteUsecase.PostAgenda(dataCore)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *AgendaDelivery) GetAgenda(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	data, err := delivery.noteUsecase.GetAgenda(idToken)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(201, helper.SuccessDataResponseHelper("success get data", fromCoreList(data)))
}

func (delivery *AgendaDelivery) GetAgendaByID(c echo.Context) error {
	id := c.Param("id")
	idNote, _ := strconv.Atoi(id)

	data, err := delivery.noteUsecase.GetAgendaByID(idNote)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(201, helper.SuccessDataResponseHelper("success get data", data))
}

func (delivery *AgendaDelivery) PutAgenda(c echo.Context) error {
	id := c.Param("id")
	idNote, _ := strconv.Atoi(id)

	var updateData AgendaRequest

	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	row, err := delivery.noteUsecase.PutAgenda(idNote, toCore(updateData))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error update data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success update data"))
}

func (delivery *AgendaDelivery) DeleteAgenda(c echo.Context) error {
	id := c.Param("id")
	idNote, _ := strconv.Atoi(id)

	row, err := delivery.noteUsecase.DeleteAgenda(idNote)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error delete data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error delete data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success delete data"))
}
