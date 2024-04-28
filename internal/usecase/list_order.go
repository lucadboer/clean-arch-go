package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

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

func (c *ListOrderUseCase) Execute() ([]entity.Order, error) {
	order, err := c.OrderRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return order, nil
}
