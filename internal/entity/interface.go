package entity

type OrderRepositoryInterface interface {
	Save(order *Order) (int64, error)
	GetTotal() (int, error)
}
