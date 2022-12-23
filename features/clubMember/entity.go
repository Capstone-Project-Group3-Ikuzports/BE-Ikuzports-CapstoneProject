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
	GetAll() (data []Core, err error)
	Create(input Core, id int) error
	GetById(id int) (data Core, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core, id int) error
	GetById(id int) (data Core, err error)
}
