package usecase

import "github.com/EuricoCruz/cleanarch_challenge/internal/entity"

type ListOrderInputDTO struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

const (
	DefaultLimit  = 50
	DefaultOffset = 0
)

type ListOrderOutputDTO struct {
	Orders []*entity.Order
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) Execute(input ListOrderInputDTO) (ListOrderOutputDTO, error) {
	if input.Limit <= 0 {
		input.Limit = DefaultLimit
	}
	if input.Offset < 0 {
		input.Offset = DefaultOffset
	}
	var Orders ListOrderOutputDTO
	orderList, err := l.OrderRepository.List(input.Offset, input.Limit)
	if err != nil {
		return Orders, err
	}
	Orders = ListOrderOutputDTO{Orders: orderList}

	return Orders, nil
}

