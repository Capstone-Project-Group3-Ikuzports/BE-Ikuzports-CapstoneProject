package clubMember

import (
	"time"
)

type Core struct {
	ID        uint
	UserID    uint
	User      User
	ClubID    uint
	Club      Club
	Status    string
	CreatedAt time.Time
	DeletedAt time.Time
}

type User struct {
	ID           uint
	Name         string
	Gender       string
	Phone_number string
	ClubMember   []Core
}

type Club struct {
	ID           uint
	Name         string
	CategoryID   uint
	Category     string
	City         string
	JoinedMember uint
	MemberTotal  uint
	Logo         string
	ClubMember   []Core
}

type ServiceInterface interface {
	GetAll(queryStatus string) (data []Core, err error)
	Create(input Core) error
	Update(input Core, id int, userId int) error
	GetById(id int) (data Core, err error)
	Delete(id int, userId int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllByStatus(queryStatus string) (data []Core, err error)
	Create(input Core) error
	Update(input Core, id int) error
	GetById(id int) (data Core, err error)
	Delete(id int) error
	FindMember(id, idUser int) (data Core, err error)
	UpdateMember(id int) error
	DeleteMember(idClub int) error
}
