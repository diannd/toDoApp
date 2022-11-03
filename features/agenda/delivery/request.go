package delivery

import (
	"toDoApp/features/agenda"
)

type AgendaRequest struct {
	UserID   uint   `json:"userid" form:"userid"`
	Title    string `json:"title" form:"title"`
	Time     string `json:"time" form:"time"`
	Desc     string `json:"desc" form:"desc"`
	Priority string `json:"priority" form:"priority"`
}

func toCore(data AgendaRequest) agenda.Core {
	return agenda.Core{
		UserID:   data.UserID,
		Title:    data.Title,
		Desc:     data.Desc,
		Priority: data.Priority,
	}
}
