package usecase

import entity "github.com/Genarodaniel/order-system/internal/entity"

type OrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	// OrderEvent      events.IEvents
	// EventDispatcher EVENT
}

func NewOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *OrderUseCase {
	return &OrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (o *OrderUseCase) CreateOrder(input CreateOrderInputDTO) (CreateOrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return CreateOrderOutputDTO{}, err
	}

	if err := order.CalculateFinalPrice(); err != nil {
		return CreateOrderOutputDTO{}, err
	}

	if err := o.OrderRepository.Save(order); err != nil {
		return CreateOrderOutputDTO{}, err
	}

	dto := CreateOrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	return dto, nil

}
