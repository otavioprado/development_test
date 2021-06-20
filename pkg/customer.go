package pkg

type Customer struct {
	Id            int
	CustomerType  string
	ArrivalTime   int
	ItemsQuantity int
	Register      Register
	NextTime      int
}

func (customer *Customer) ChooseRegister(time int, store *GroceryStore) {
	var register *Register
	switch ct := customer.CustomerType; ct {
	case "A":
		register = store.getShortestLine()
		customer.Register = *register
	case "B":
		register = store.getFewestNumberOfItemsLine()
		customer.Register = *register
	default:
		panic("invalid customer type")
	}

	if customer.Register.InTraining {
		customer.NextTime = time + 2
	} else {
		customer.NextTime = time + 1
	}

	store.Registers[register.Id].CustomersQueue = append(store.Registers[register.Id].CustomersQueue, *customer)
}

