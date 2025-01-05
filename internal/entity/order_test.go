package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	t.Run("NewOrder - Should return error invalid ID", func(t *testing.T) {
		o := Order{}
		assert.EqualError(t, o.IsValid(), "Invalid ID")
	})

	t.Run("NewOrder - Should return error invalid Price", func(t *testing.T) {
		o := Order{
			ID: uuid.NewString(),
		}

		assert.EqualError(t, o.IsValid(), "Invalid Price")
	})

	t.Run("NewOrder - Should return error Invalid Tax", func(t *testing.T) {
		o := Order{
			ID:    uuid.NewString(),
			Price: 3.50,
		}

		assert.EqualError(t, o.IsValid(), "Invalid Tax")
	})

	t.Run("NewOrder - Should return success", func(t *testing.T) {
		price := 3.50
		tax := 2.0

		newOrder, err := NewOrder(uuid.NewString(), price, tax)

		assert.Nil(t, err)
		assert.Nil(t, newOrder.IsValid())
		assert.Equal(t, price, newOrder.Price)
		assert.Equal(t, tax, newOrder.Tax)
	})

}

func TestCalculatePrices(t *testing.T) {

	t.Run("CalculatePrices - Should return the price calculated", func(t *testing.T) {
		price := 3.50
		tax := 2.0

		newOrder, err := NewOrder(uuid.NewString(), price, tax)

		assert.Nil(t, err)
		assert.Nil(t, newOrder.IsValid())
		assert.Nil(t, newOrder.CalculateFinalPrice())
		assert.Equal(t, price+tax, newOrder.FinalPrice)
	})

	t.Run("CalculatePrices - Should return error when the struct used to calculate is invalid", func(t *testing.T) {
		price := 3.50
		tax := 2.0
		order := Order{
			Price: price,
			Tax:   tax,
		}
		assert.EqualError(t, order.CalculateFinalPrice(), "Invalid ID")
	})

}
