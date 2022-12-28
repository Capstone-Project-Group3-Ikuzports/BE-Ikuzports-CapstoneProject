package transaction

type TransactionCore struct {
	ID              uint
	UserID          int
	TotalPrice      int
	TotalQuantity   int
	ProductID       uint
	TransactionID   string
	StatusPayment   string
	VirtualAccount  string
	TransactionTime string
	OrderID         string
}

type MidtransCore struct {
	OrderID       string
	GrossAmt      string
	StatusMessage string
	VANumbers     VANumbersCore
}

type VANumbersCore struct {
	Bank     string
	VANumber string
}

type RepositoryInterface interface {
	Create(input TransactionCore) (row int, err error)
	GetAll() (data []TransactionCore, err error)
	GetByID(id int) (data TransactionCore, err error)
	Update(input TransactionCore) (rows int, err error)
}

type ServiceInterface interface {
	Create(input TransactionCore) (data MidtransCore, err error)
	GetAll() (data []TransactionCore, err error)
	GetByID(id int) (data TransactionCore, err error)
	Update(input TransactionCore) (err error)
}
