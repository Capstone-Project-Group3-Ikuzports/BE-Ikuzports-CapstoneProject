package product

type ProductCore struct {
	ID             uint
	Name           string `validate:"required"`
	Price          uint   `validate:"required"`
	Description    string `validate:"required"`
	UserID         int
	User           User
	ItemCategoryID uint `validate:"required"`
	ItemCategory   ItemCategory
	City           string
	ProductImage   []ProductImage
}

type ItemCategory struct {
	ID   uint
	Name string
}

type User struct {
	ID   uint
	Name string
	City string
}

type ProductImage struct {
	ID  uint
	URL string
}

type RepositoryInterface interface {
	GetAll() (data []ProductCore, err error)
	GetAllFilter(queryItemCategoryID int, queryCity, queryName string) (data []ProductCore, err error)
	Create(input ProductCore) (row int, err error)
	GetByID(id int) (dataCore ProductCore, err error)
	Update(id int, input ProductCore) (row int, err error)
	Delete(id int) (row int, err error)
	GetImages(id int) (data []ProductImage, err error)
}

type ServiceInterface interface {
	GetAll(queryItemCategoryID int, queryCity, queryName string) (data []ProductCore, err error)
	Create(input ProductCore) (err error)
	GetByID(id int) (dataCore ProductCore, err error)
	Update(id int, input ProductCore) (err error)
	Delete(id int) (err error)
	GetImages(id int) (data []ProductImage, err error)
}
