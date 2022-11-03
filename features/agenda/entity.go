package agenda

import "time"

type Core struct {
	ID       uint
	UserID   uint
	Title    string
	Time     time.Time
	Desc     string
	Priority string
}

type UsecaseInterface interface {
	PostAgenda(Core) (int, error)
	GetAgenda(token int) ([]Core, error)
	GetAgendaByID(id int) (Core, error)
	PutAgenda(id int, data Core) (int, error)
	DeleteAgenda(id int) (int, error)
}

type DataInterface interface {
	CreateAgenda(Core) (int, error)
	ReadAgenda(token int) ([]Core, error)
	ReadAgendaByID(id int) (Core, error)
	UpdateAgenda(id int, data Core) (int, error)
	DeleteData(id int) (int, error)
}
