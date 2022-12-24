package category

type CategoryCore struct {
	ID   uint
	Name string
}

type RepositoryInterface interface {
	GetAll() (data []CategoryCore, err error)
}

type ServiceInterface interface {
	GetAll() (data []CategoryCore, err error)
}
