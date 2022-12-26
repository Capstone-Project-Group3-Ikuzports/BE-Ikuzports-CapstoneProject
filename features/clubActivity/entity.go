package clubActivity

type Core struct {
	ID             uint
	ClubID         uint
	Name           string `validate:"required"`
	StartTime      string `validate:"required"`
	EndTime        string `validate:"required"`
	Day            string `validate:"required"`
	Location       string `validate:"required"`
	ActivityDetail string
}

type ServiceInterface interface {
	GetAll(queryClubId int) (data []Core, err error)
	Create(input Core, idUser int) error
	Update(input Core, id int, userId int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllByClubId(queryClubId uint) (data []Core, err error)
	Create(input Core, idUser int) error
	Update(input Core, id int) error
}
