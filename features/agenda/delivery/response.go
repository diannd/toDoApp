package delivery

import (
	"time"
	"toDoApp/features/agenda"
)

type AgendaResponse struct {
	ID       uint      `json:"id" form:"id"`
	UserID   uint      `json:"userid" form:"userid"`
	Title    string    `json:"title" form:"title"`
	Time     time.Time `json:"time" form:"time"`
	Desc     string    `json:"desc" form:"desc"`
	Priority string    `json:"priority" form:"priority"`
}

func fromCore(data agenda.Core) AgendaResponse {
	return AgendaResponse{
		ID:       data.ID,
		UserID:   data.UserID,
		Title:    data.Title,
		Time:     data.Time,
		Desc:     data.Desc,
		Priority: data.Priority,
	}
}

func fromCoreList(data []agenda.Core) []AgendaResponse {
	var dataRes []AgendaResponse
	for _, v := range data {
		dataRes = append(dataRes, AgendaResponse{
			ID:       v.ID,
			UserID:   v.UserID,
			Title:    v.Title,
			Time:     v.Time,
			Desc:     v.Desc,
			Priority: v.Priority,
		})
	}
	return dataRes
}
