package pkg

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var successTests = []struct {
	path           string
	expectedResult int
}{
	{"../input1.txt", 7},
	{"../input2.txt", 13},
	{"../input3.txt", 6},
	{"../input4.txt", 9},
	{"../input5.txt", 11},
	{"../input8.txt", 50},
}

func TestInputs(t *testing.T) {
	for _, e := range successTests {
		file, _ := ioutil.ReadFile(e.path)
		store, _ := Setup(file)
		result, _ := StartSimulation(store)
		assert.Equal(t, e.expectedResult, result)
	}
}

func TestInvalidCustomerType(t *testing.T) {
	file, _ := ioutil.ReadFile("../input6.txt")
	store, _ := Setup(file)
	_, err := StartSimulation(store)
	assert.EqualError(t, err, "invalid customer type")
}

func TestInvalidRegistersNumber(t *testing.T) {
	file, _ := ioutil.ReadFile("../input7.txt")
	_, err := Setup(file)
	assert.EqualError(t, err, "invalid number of registers")
}

func TestSetup(t *testing.T) {
	file, _ := ioutil.ReadFile("../input1.txt")
	setup, _ := Setup(file)

	assert.NotNil(t, setup)
	assert.Equal(t, 1, len(setup.Registers))
	assert.Equal(t, 0, len(setup.Registers[1].CustomersQueue))
	assert.True(t, setup.Registers[1].InTraining)

	assert.Equal(t, 2, len(setup.Customers))
	assert.Equal(t, 1, len(setup.Customers[1]))
	assert.Equal(t, 1, len(setup.Customers[2]))

	assert.Equal(t, 2, setup.Customers[1][0].ItemsQuantity)
	assert.Equal(t, 1, setup.Customers[2][0].ItemsQuantity)

	assert.Equal(t, "A", setup.Customers[1][0].CustomerType)
	assert.Equal(t, "A", setup.Customers[2][0].CustomerType)

	assert.Equal(t, 1, setup.NumberOfRegisters)
}
