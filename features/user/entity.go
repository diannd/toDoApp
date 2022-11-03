package user

type Core struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

type UsecaseInterface interface {
	PostUser(Core) (int, error)
	GetUser(id int) (Core, error)
	PutUser(id int, data Core) (int, error)
	DeleteUser(id int) (int, error)
}

type DataInterface interface {
	RegisterUser(Core) (int, error)
	ReadUser(id int) (Core, error)
	UpdateUser(id int, data Core) (int, error)
	DeleteData(id int) (int, error)
}
