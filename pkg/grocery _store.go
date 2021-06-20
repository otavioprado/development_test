package pkg

type GroceryStore struct {
	NumberOfRegisters int
	Customers         map[int][]Customer
	Registers         map[int]*Register
	TotalCustomers    int
}

func newGroceryStore() *GroceryStore {
	store := GroceryStore{
		NumberOfRegisters: 0,
		Customers:         make(map[int][]Customer),
		Registers:         make(map[int]*Register),
	}
	return &store
}

func (groceryStore *GroceryStore) isStoreInService() bool {
	if groceryStore.TotalCustomers > 0 {
		return true
	}
	return false
}

func (groceryStore *GroceryStore) ProcessOneMoreItem(currentTimeSecond int) {
	for _, register := range groceryStore.Registers {

		if len(register.CustomersQueue) != 0 && register.CustomersQueue[0].NextTime == currentTimeSecond {
			register.CustomersQueue[0].ItemsQuantity--

			if register.CustomersQueue[0].ItemsQuantity <= 0 {
				register.CustomersQueue = register.CustomersQueue[1:]
				groceryStore.TotalCustomers--
				if len(register.CustomersQueue) == 0 {
					continue
				}
			}

			if register.InTraining {
				register.CustomersQueue[0].NextTime = currentTimeSecond + 2
			} else {
				register.CustomersQueue[0].NextTime = currentTimeSecond + 1
			}
		}
	}
}

func (groceryStore *GroceryStore) getFewestNumberOfItemsLine() *Register {
	fewest := groceryStore.Registers[1]
	for _, register := range groceryStore.Registers {
		if len(register.CustomersQueue) == 0 {
			return register
		}

		lastIdFewest := len(fewest.CustomersQueue) - 1
		lastIdRegister := len(register.CustomersQueue) - 1

		if register.CustomersQueue[lastIdRegister].ItemsQuantity < fewest.CustomersQueue[lastIdFewest].ItemsQuantity {
			fewest = register
		}
	}
	return fewest
}

func (groceryStore *GroceryStore) getShortestLine() *Register {
	shortest := groceryStore.Registers[1]
	for _, register := range groceryStore.Registers {
		if len(register.CustomersQueue) < len(shortest.CustomersQueue) {
			shortest = register
		}
	}
	return shortest
}
