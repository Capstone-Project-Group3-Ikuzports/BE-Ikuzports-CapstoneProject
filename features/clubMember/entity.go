package clubMember

import "time"

type Core struct {
	ID        uint
	UserID    uint
	ClubID    uint
	Status    string
	CreatedAt time.Time
}

type ServiceInterface interface {
	GetAll(queryStatus string) (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllByStatus(queryStatus string) (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	FindMember(id, idUser int) (data Core, err error)
	UpdateMember(id int) error
}
