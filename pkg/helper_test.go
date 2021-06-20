package pkg

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNewGroceryStore(t *testing.T) {
	store := newGroceryStore()

	assert.NotNil(t, store)
	assert.Equal(t, 0, store.NumberOfRegisters)
	assert.Equal(t, 0, len(store.Customers))
	assert.Equal(t, 0, len(store.Registers))
}

func TestSetup(t *testing.T) {
	file, err := ioutil.ReadFile("../input1.txt")
	assert.Nil(t, err)

	setup := Setup(file)
	assert.NotNil(t, setup)

	assert.Equal(t, 1, len(setup.Registers))
	assert.True(t, setup.Registers[1].InTraining)
	assert.Equal(t, 0, len(setup.Registers[1].CustomersQueue))

	assert.Equal(t, 2, len(setup.Customers))
	assert.Equal(t, 1, len(setup.Customers[1]))

	assert.Equal(t, "A", setup.Customers[1][0].CustomerType)
	assert.Equal(t, 2, setup.Customers[1][0].ItemsQuantity)

	assert.Equal(t, "A", setup.Customers[2][0].CustomerType)
	assert.Equal(t, 1, setup.Customers[2][0].ItemsQuantity)

	assert.Equal(t, 1, setup.NumberOfRegisters)
}

func TestInput1(t *testing.T) {
	file, err := ioutil.ReadFile("../input1.txt")
	assert.Nil(t, err)
	assert.Equal(t, 7, StartSimulation(Setup(file)))
}

func TestInput2(t *testing.T) {
	file, err := ioutil.ReadFile("../input2.txt")
	assert.Nil(t, err)
	assert.Equal(t, 13, StartSimulation(Setup(file)))
}

func TestInput3(t *testing.T) {
	file, err := ioutil.ReadFile("../input3.txt")
	assert.Nil(t, err)
	assert.Equal(t, 6, StartSimulation(Setup(file)))
}

func TestInput4(t *testing.T) {
	file, err := ioutil.ReadFile("../input4.txt")
	assert.Nil(t, err)
	assert.Equal(t, 9, StartSimulation(Setup(file)))
}

func TestInput5(t *testing.T) {
	file, err := ioutil.ReadFile("../input5.txt")
	assert.Nil(t, err)
	assert.Equal(t, 11, StartSimulation(Setup(file)))
}