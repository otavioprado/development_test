package pkg

import (
	"sort"
	"strconv"
	"strings"
)

type GroceryStore struct {
	NumberOfRegisters int
	Customers         map[int][]Customer
	Registers         map[int]*Register
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
	for _, register := range groceryStore.Registers {
		if len(register.CustomersQueue) != 0 {
			return true
		}
	}
	return false
}

func (groceryStore *GroceryStore) ProcessOneMoreItem(currentTimeSecond int) {
	for _, register := range groceryStore.Registers {

		if len(register.CustomersQueue) != 0 && register.CustomersQueue[0].NextTime == currentTimeSecond {
			register.CustomersQueue[0].ItemsQuantity--

			if register.CustomersQueue[0].ItemsQuantity <= 0 {
				register.CustomersQueue = register.CustomersQueue[1:]
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

func Setup(file []byte) *GroceryStore {
	store := newGroceryStore()

	dataSlice := strings.Split(string(file), "\n")
	for i, line := range dataSlice {
		if i == 0 {
			store.NumberOfRegisters = toInt(dataSlice[i])
			for y := 1; y <= store.NumberOfRegisters; y++ {
				r := Register{
					Id:         y,
					InTraining: y == store.NumberOfRegisters,
				}
				store.Registers[y] = &r
			}
		} else {
			splitLine := strings.Split(line, " ")
			c := Customer{
				Id:            i,
				CustomerType:  splitLine[0],
				ArrivalTime:   toInt(splitLine[1]),
				ItemsQuantity: toInt(splitLine[2]),
			}

			store.Customers[c.ArrivalTime] = append(store.Customers[c.ArrivalTime], c)
		}
	}

	// If two or more customers arrive at the same time, those with fewer items choose registers before those with more items.
	for _, customer := range store.Customers {
		sort.Slice(customer, func(i, j int) bool {
			if customer[i].ItemsQuantity == customer[j].ItemsQuantity {
				return customer[i].CustomerType == "A"
			}

			return customer[i].ItemsQuantity < customer[j].ItemsQuantity
		})
	}

	return store
}

func toInt(number string) int {
	parseInt, err := strconv.ParseInt(number, 10, 64)
	check(err)
	return int(parseInt)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StartSimulation(store *GroceryStore) int {
	time := 1

	for {
		store.ProcessOneMoreItem(time)

		// 8) If two or more customers arrive at the same time, those with fewer items choose registers before those with more items. If the customers have the same number of items, then type A customers choose before type B customers.
		if customers, found := store.Customers[time]; found {
			for _, customer := range customers {
				customer.ChooseRegister(time, store)
			}
		}

		if !store.isStoreInService() {
			break
		}

		time++
	}

	return time
}
