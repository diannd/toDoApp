package note

type Core struct {
	ID     uint
	UserID uint
	Title  string
	Note   string
}

type UsecaseInterface interface {
	PostNote(data Core) (int, error)
	GetNoteByID(id int) (Core, error)
	GetNote(token int) ([]Core, error)
	PutNote(id int, data Core) (int, error)
	DeleteNote(id int) (int, error)
}

type DataInterface interface {
	CreateNote(data Core) (int, error)
	ReadNoteByID(id int) (Core, error)
	ReadNote(token int) ([]Core, error)
	UpdateNote(id int, data Core) (int, error)
	DeleteData(id int) (int, error)
}
