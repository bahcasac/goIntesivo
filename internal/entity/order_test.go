package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateFinalPrice(t *testing.T) {
	t.Run("it should return an error when ID is blank", func(t *testing.T) {
		order := Order{}
		assert.Error(t, order.Validate(), "invalid id")
	})

	t.Run("it should return an error when price is blank", func(t *testing.T) {
		order := Order{ID: "123"}
		assert.Error(t, order.Validate(), "invalid price")
	})

	t.Run("it should return an error when tax is blank", func(t *testing.T) {
		order := Order{ID: "123", Price: 100}
		assert.Error(t, order.Validate(), "invalid tax")
	})

	t.Run("it should return a order successfully", func(t *testing.T) {
		order := Order{ID: "123", Price: 10.0, Tax: 10.0}

		assert.NoError(t, order.Validate())
		assert.Equal(t, 10.0, order.Price)
		assert.Equal(t, 10.0, order.Tax)

		err := order.CalculateFinalPrice()
		assert.NoError(t, err)
		assert.Equal(t, 20.0, order.FinalPrice)
	})
}
