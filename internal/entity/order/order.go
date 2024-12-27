package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Order struct {
	ID         string
	Price      float64 `validate:"required"`
	FinalPrice float64
	Tax        float64 `validate:"required"`
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	if err := order.IsValid(); err != nil {
		return nil, err
	}

	if err := order.CalculateFinalPrice(); err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) IsValid() error {
	if err := uuid.Validate(o.ID); err != nil {
		return errors.New("Invalid ID")
	}

	if o.Price <= 0 {
		return errors.New("Invalid Price")
	}

	if o.Tax <= 0 {
		return errors.New("Invalid Tax")
	}

	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	if err := o.IsValid(); err != nil {
		return err
	}

	return nil
}
