package pkg

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

func Setup(file []byte) (*GroceryStore, error) {
	store := newGroceryStore()

	dataSlice := strings.Split(string(file), "\n")
	for i, line := range dataSlice {
		if i == 0 {
			store.NumberOfRegisters = toInt(dataSlice[i])
			if store.NumberOfRegisters <= 0 {
				return nil, errors.New("invalid number of registers")
			}

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
			store.TotalCustomers++

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

	return store, nil
}

func toInt(number string) int {
	parseInt, err := strconv.ParseInt(number, 10, 64)
	Check(err)
	return int(parseInt)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func StartSimulation(store *GroceryStore) (int, error) {
	time := 1

	for {
		store.ProcessOneMoreItem(time)
		if customers, found := store.Customers[time]; found {
			for _, customer := range customers {
				err := customer.ChooseRegister(time, store)
				if err != nil {
					return 0, err
				}
			}
		}

		if !store.isStoreInService() {
			break
		}

		time++
	}

	return time, nil
}
