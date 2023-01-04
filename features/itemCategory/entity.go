package itemCategory

type ItemCategoryCore struct {
	ID   uint
	Name string
}

type RepositoryInterface interface {
	GetAll() (data []ItemCategoryCore, err error)
}

type ServiceInterface interface {
	GetAll() (data []ItemCategoryCore, err error)
}
