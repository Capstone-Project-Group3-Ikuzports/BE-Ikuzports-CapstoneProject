package club

import (
	"ikuzports/features/chat"
	"ikuzports/features/clubActivity"
	_members "ikuzports/features/clubMember"
	"ikuzports/features/galery"
	"time"
)

type Core struct {
	ID           uint
	Name         string `validate:"required"`
	Address      string
	City         string `validate:"required"`
	CategoryID   uint   `validate:"required"`
	Description  string
	Logo         string
	JoinedMember uint
	MemberTotal  uint `validate:"required"`
	Rule         string
	Requirement  string `validate:"required"`
	Category     Category
	Member       []ClubMember
	CreatedAt    time.Time
	UpdateAt     time.Time
}

type Category struct {
	ID   uint
	Name string
}

type User struct {
	ID   uint
	Name string
}
type Status struct {
	ID     uint
	UserID uint
	ClubID uint
	Status string
}

type Chat struct {
	chat.Core
}

type Galery struct {
	galery.Core
}

type ClubMember struct {
	ID     uint
	UserID uint
	ClubID uint
	Status string
}
type ServiceInterface interface {
	GetAll(queryName string, queryCity string, queryCategoryID int) (data []Core, err error)
	Create(input Core, id int) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int, userId int) error
	Delete(id int, userId int) error
	GetMembers(id int) (data []_members.Core, err error)
	GetChats(id int) (data []chat.Core, err error)
	GetGaleries(id int) (data []galery.Core, err error)
	GetActivities(id int) (data []clubActivity.Core, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(queryName string, queryCity string, queryCategoryID int) (data []Core, err error)
	Create(input Core, id int) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	GetLastID() (id int, err error)
	UpdateMember(id int) (rows int, err error)
	GetMembers(id int) (data []_members.Core, err error)
	GetChats(id int) (data []chat.Core, err error)
	GetGaleries(id int) (data []galery.Core, err error)
	GetActivities(id int) (data []clubActivity.Core, err error)
	GetStatus(id int, userId int) (data Status, err error)
}
