package delivery

import (
	"strconv"
	"toDoApp/features/note"
	"toDoApp/middlewares"
	"toDoApp/utils/helper"

	"github.com/labstack/echo/v4"
)

type NoteDelivery struct {
	noteUsecase note.UsecaseInterface
}

func New(e *echo.Echo, usecase note.UsecaseInterface) {
	handler := &NoteDelivery{
		noteUsecase: usecase,
	}

	e.POST("/note", handler.PostNote, middlewares.JWTMiddleware())
	e.GET("/note/:id", handler.GetNoteByID, middlewares.JWTMiddleware())
	e.GET("/note", handler.GetNote, middlewares.JWTMiddleware())
	e.PUT("/note/:id", handler.PutNote, middlewares.JWTMiddleware())
	e.DELETE("/note/:id", handler.DeleteNote, middlewares.JWTMiddleware())
}

func (delivery *NoteDelivery) PostNote(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	var data NoteRequest

	errBind := c.Bind(&data)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error binda data"))
	}

	data.UserID = uint(idToken)

	row, err := delivery.noteUsecase.PostNote(toCore(data))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *NoteDelivery) GetNoteByID(c echo.Context) error {
	id := c.Param("id")
	idNote, _ := strconv.Atoi(id)

	data, err := delivery.noteUsecase.GetNoteByID(idNote)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(201, helper.SuccessDataResponseHelper("success get data", data))
}

func (delivery *NoteDelivery) GetNote(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	data, err := delivery.noteUsecase.GetNote(idToken)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(201, helper.SuccessDataResponseHelper("success get data", fromCoreList(data)))
}

func (delivery *NoteDelivery) PutNote(c echo.Context) error {
	id := c.Param("id")
	idNote, _ := strconv.Atoi(id)

	var updateData NoteRequest

	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	row, err := delivery.noteUsecase.PutNote(idNote, toCore(updateData))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error update data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success update data"))
}

func (delivery *NoteDelivery) DeleteNote(c echo.Context) error {
	id := c.Param("id")
	idNote, _ := strconv.Atoi(id)

	row, err := delivery.noteUsecase.DeleteNote(idNote)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error delete data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error delete data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success delete data"))
}
