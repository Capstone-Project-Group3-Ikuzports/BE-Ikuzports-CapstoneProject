package productImage

type ProductImageCore struct {
	ID        uint
	URL       string
	ProductID uint
}

type RepositoryInterface interface {
	Create(data ProductImageCore) (rows int, err error)
	GetAll() (data []ProductImageCore, err error)
	GetByID(id int) (data ProductImageCore, err error)
}

type ServiceInterface interface {
	Create(data ProductImageCore) (err error)
	GetAll() (data []ProductImageCore, err error)
	GetByID(id int) (data ProductImageCore, err error)
}
