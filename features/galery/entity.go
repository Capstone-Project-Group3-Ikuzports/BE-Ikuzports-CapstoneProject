package galery

type Core struct {
	ID      uint
	Url     string `validate:"required"`
	ClubID  uint
	Caption string `validate:"required"`
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core, id int) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int, userId int) error
	Delete(id int, userId int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core, id int) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	// GetLastID() (id int, err error)
	// UpdateMember(id int) (rows int, err error)
	// GetStatus(id int, userId int) (data Status, err error)
}
