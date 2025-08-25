package usecase

import "github.com/EuricoCruz/cleanarch_challenge/internal/entity"

type ListOrderInputDTO struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

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
	var Orders ListOrderOutputDTO
	orderList, err := l.OrderRepository.List(input.Limit, input.Offset)
	if err != nil {
		return Orders, err
	}
	Orders = ListOrderOutputDTO{Orders: orderList}

	return Orders, nil
}

