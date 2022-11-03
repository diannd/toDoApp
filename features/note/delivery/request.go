package delivery

import "toDoApp/features/note"

type NoteRequest struct {
	UserID uint   `json:"userid" form:"userid"`
	Title  string `json:"title" form:"title"`
	Note   string `json:"note" form:"note"`
}

func toCore(data NoteRequest) note.Core {
	return note.Core{
		UserID: data.UserID,
		Title:  data.Title,
		Note:   data.Note,
	}
}
