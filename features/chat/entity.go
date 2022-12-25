package chat

import "time"

type Core struct {
	ID        uint
	UserID    uint
	User      User
	ClubID    uint
	Message   string
	CreatedAt time.Time
}

type User struct {
	ID   uint
	Name string
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	Delete(id int, userId int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetById(id int) (data Core, err error)
	Create(input Core) error
	Delete(id int) error
	FindMember(id, idUser int) (data Core, err error)
}
