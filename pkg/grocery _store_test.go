package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGroceryStore(t *testing.T) {
	store := newGroceryStore()
	assert.NotNil(t, store)
	assert.Equal(t, 0, store.NumberOfRegisters)
	assert.Equal(t, 0, len(store.Customers))
	assert.Equal(t, 0, len(store.Registers))
}
