package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotal() (int, error)
	List(offset int, limit int) ([]*Order, error)
}

