package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// de forma nativa
func TestIfItGetsAnErrorIfIDisBlankNative(t *testing.T) {
	order:= Order{}
	if order.Validate() == nil	{
		t.Error("ID is required")
	}
}
func TestIfItGetsAnErrorIfIDisBlank(t *testing.T) {
	order:= Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestFinalPrice( t *testing.T)  {
	order:= Order{
		ID: "123", Price: 10.0, Tax: 1.2,
	}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.2, order.Tax)

	order.CalculateFinalPrice()
	assert.Equal(t, 11.2, order.FinalPrice)
}