package delivery

import "toDoApp/features/note"

type NoteResponse struct {
	ID     uint   `json:"id" form:"id"`
	UserID uint   `json:"userid" form:"userid"`
	Title  string `json:"title" form:"title"`
	Note   string `json:"note" form:"note"`
}

func fromCore(data note.Core) NoteResponse {
	return NoteResponse{
		ID:     data.ID,
		UserID: data.UserID,
		Title:  data.Title,
		Note:   data.Note,
	}
}

func fromCoreList(data []note.Core) []NoteResponse {
	var dataRes []NoteResponse
	for _, v := range data {
		dataRes = append(dataRes, NoteResponse{
			ID:     v.ID,
			UserID: v.UserID,
			Title:  v.Title,
			Note:   v.Note,
		})
	}
	return dataRes
}
