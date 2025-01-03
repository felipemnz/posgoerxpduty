package usecase

import "github.com/felipemnz/go-expert-challenge-cleanarch/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: OrderRepository}
}

func (c *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()

	if err != nil {
		return []OrderOutputDTO{}, err
	}

	var ordersOutputDTO []OrderOutputDTO
	for _, order := range orders {
		ordersOutputDTO = append(ordersOutputDTO, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ordersOutputDTO, nil
}
